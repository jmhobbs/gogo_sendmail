[![Build Status](https://travis-ci.org/jmhobbs/gogo_sendmail.svg?branch=master)](https://travis-ci.org/jmhobbs/gogo_sendmail)

# GoGo Sendmail

This is a golang clone of [mini_sendmail](http://acme.com/software/mini_sendmail/).

It's hopefully easier to read, safer and can be enhanced.

# Differences

Since I'm using go's [net/smtp.SendMail](https://golang.org/pkg/net/smtp/#SendMail) I don't really have anything to print in verbose mode.

I also haven't wired the timeouts up, for the same reasons.

There's probably some other differences too, I was only referring to the mini_sendmail code occasionally and intuiting the other behavior.

# Usage

You should be able to use it as a drop in.

```
Usage of gogo_sendmail:
  -T int
    	Specifies timeout - defaults to one minute. (default 60)
  -f string
    	Sets the name of the "from" person (i.e. the sender of the mail).
  -p int
    	Specifies the port to use.  Without this it uses 25, the standard SMTP port. (default 25)
  -s string
    	Specifies the SMTP server to use.  Without this it uses localhost. (default "localhost")
  -t	Read message for recipients.  To:, Cc:, and Bcc: lines will be scanned for recipient addresses.  The Bcc: line will be deleted before transmission.
  -v	Verbose mode - shows the conversation with the SMTP server.
```
