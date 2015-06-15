# ec2s

EC2 instance viewer

WARNING: This project is under development.

# Configure

```
$ cat ~/.ec2s
[aws]
ACCESS_KEY_ID=xxxx
SECRET_ACCESS_KEY=yyyy
REGION=ap-northeast-1
```

# Usage

## list subcommand

```
# show ec2 instances
$ ec2s list

# show ec2 instances in the vpc
$ ec2s list --vpc 'name'

$ ec2s list --vpcid vpc-xxxx
```

## vpcs subcommand

```
# show vpcs
$ ec2s vpcs
```

## License

MIT License

## Contributing

1. Fork it ( https://github.com/[my-github-username]/ec2_meta/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

