package main

import "github.com/codegangsta/cli"

func BoxNameFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "boxname, b",
		Usage: "name of the box",
	}
}

func BoxFileNameFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "boxfile, f",
		Usage: "path to the box file",
	}
}

func BoxRmFlag() cli.BoolFlag {
	return cli.BoolFlag{
		Name:  "remove, r",
		Usage: "remove box file when deleting version",
	}
}

func BoxVerFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "version, v",
		Usage: "box version",
	}
}

func ProviderFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "provider, p",
		Usage: "provider of the version",
	}
}

func QuietFlag() cli.BoolFlag {
	return cli.BoolFlag{
		Name:  "quiet, q",
		Usage: "suppress output to stdout ",
	}
}

func NoopFlag() cli.BoolFlag {
	return cli.BoolFlag{
		Name:  "noop, n",
		Usage: "run in no-op mode",
	}
}

func DescFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "description, d",
		Usage: "description",
	}
}

func DescShortFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "shortdescription, s",
		Usage: "short box description",
	}
}

func InputFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "input, i",
		Usage: "file to read metadata from",
	}
}

func OutputFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:  "output, o",
		Usage: "file to write metadata to",
	}
}
