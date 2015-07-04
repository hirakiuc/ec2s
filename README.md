# ec2s

EC2 instance operation tool with ssh, peco.

WARNING: This project is under development.

# Configure

```
$ cat ~/.ec2s
[aws]
AWS_ACCESS_KEY_ID="xxxx"
AWS_SECRET_ACCESS_KEY="yyyy"
AWS_REGION="ap-northeast-1"

[peco]
path = "/path/to/peco"

[ssh]
port = 22
user = "my-user"

[[ssh.identity_file]]
name = "my_keypair"
path = "~/.ssh/my_keypair.private_key"

[[ssh.identity_file]]
name = "my_other_keypair"
path = "~/.ssh/my_other_keypair.private_key"
```

# Usage

## list subcommand

```
# show ec2 instances
$ ec2s list

# show ec2 instances in the vpc
$ ec2s list -vpc-id vpc-xxxx

$ ec2s list -vpc-name vpcname
```

## ssh subcommand

```
# login via ssh to the ec2 instance.
$ ec2s ssh

$ ec2s ssh -vpc-id vpc-xxxx

$ ec2s ssh -vpc-name vpcname
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

