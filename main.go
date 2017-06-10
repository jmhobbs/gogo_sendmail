package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"os/user"
	"strings"
)

type options struct {
	verbose            bool
	from               string
	extract_recipients bool
	server             string
	port               int
	timeout            int
}

func main() {
	opts := options{}
	// TODO: What to print?
	flag.BoolVar(&opts.verbose, "v", false, "Verbose mode - shows the conversation with the SMTP server.")
	flag.StringVar(&opts.from, "f", "", "Sets the name of the \"from\" person (i.e. the sender of the mail).")
	flag.BoolVar(&opts.extract_recipients, "t", false, "Read message for recipients.  To:, Cc:, and Bcc: lines will be scanned for recipient addresses.  The Bcc: line will be deleted before transmission.")
	flag.StringVar(&opts.server, "s", "localhost", "Specifies the SMTP server to use.  Without this it uses localhost.")
	flag.IntVar(&opts.port, "p", 25, "Specifies the port to use.  Without this it uses 25, the standard SMTP port.")
	// TODO: Handle timeouts
	flag.IntVar(&opts.timeout, "T", 60, "Specifies timeout - defaults to one minute.")
	flag.Parse()

	recipients := flag.Args()

	// Default to <user>@<hostname> if no From provided
	// mini_sendmail.c L177
	if opts.from == "" {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		opts.from = user.Username
	}

	if !strings.Contains(opts.from, "@") {
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}

		opts.from = fmt.Sprintf("%s@%s", opts.from, hostname)
	}

	headers, body := readEmail(os.Stdin)

	if opts.extract_recipients {
		var err error
		// TODO: Merge with command line recipients?
		recipients, err = extractRecipients(strings.Join(append(headers, "", "Body"), "\r\n"))
		if err != nil {
			log.Fatalf("extracting recipients: %s\n", err)
		}

		// Strip Bcc
		// mini_sendmail.c L521
		for i, header := range headers {
			if strings.ToLower(header[:4]) == "bcc:" {
				headers = append(headers[:i], headers[i+1:]...)
				break
			}
		}
	}

	err := smtp.SendMail(fmt.Sprintf("%s:%d", opts.server, opts.port), nil, opts.from, recipients, []byte(strings.Join(append(headers, body...), "\r\n")))
	if err != nil {
		log.Fatal(err)
	}
}

func readEmail(src io.Reader) (headers, body []string) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		headers = append(headers, line)
	}

	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	return headers, body
}

func extractRecipients(headers string) ([]string, error) {
	r := strings.NewReader(headers)
	m, err := mail.ReadMessage(r)
	if err != nil {
		return []string{}, err
	}

	var (
		addrs []*mail.Address
		none  struct{}
	)

	recipients := make(map[string]struct{})

	for _, header := range []string{"To", "Cc", "Bcc"} {
		addrs, err = m.Header.AddressList(header)
		if err != nil && err != mail.ErrHeaderNotPresent {
			return []string{}, err
		}

		for _, addr := range addrs {
			recipients[addr.Address] = none
		}
	}

	list := make([]string, 0, len(recipients))
	for email, _ := range recipients {
		list = append(list, email)
	}

	return list, nil
}
