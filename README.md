# EC2 Metadata Dump

EC2 metadata dump is a tool to grab the data from the [EC2 metadata service](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html) and output it as json.

## Quick Start

First, download a pre-built binary from the desired [release](https://github.com/thbishop/ec2_metadata_dump/releases).

You can then dump the data with:

```
$ ec2_metadata_dump
...
```

The data will be sent to `STDOUT` and errors will be sent to `STDERR`.

## Support
Today, there is support for:

* dumping data only from the `latest` endpoint (http://169.254.169.254/latest/)
* dumping `meta-data` and `user-data` (the `dynamic` portion is *not* dumped)

## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Send a pull request and I'll get it integrated

## License
See LICENSE
