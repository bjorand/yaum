provider "aws" {
  region = "eu-west-1"
}

resource "aws_key_pair" "bjo" {
  key_name   = "yaum-bjo"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDU+QaY8Tne/znhTCvlpyPoxtkDVVXqXbEzaGTATFBmgj/L67yf358FlvsGaqAPGlJxECdwAG/4yd7/OT1yFS561hRf2vg2Jrps6xgttafZ2ddZDovY7MxlsZUZfOY+BsVa1GjIMtHTzY7eVEereWmMPlU8NKGzLr9AFJuCK9uTd8tpFglpmK/vk9AIpLFnz1TFkbfeANe3V8YqhCxkbaMhijLqxUWXz1vxEiwpwriOLdBx+EdGqXFcGwtMqP1QgCa7v2u9/2ye5R/qHlrVOMAk4n/qiF8yRNIVs1ZLgNn52tFnW0KuwaGr2sUJgds2kkXzZOoAS3WqA7xxrHMriGVv"
}


resource "aws_vpc" "yaum" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
}

resource "aws_internet_gateway" "yaum" {
  vpc_id = "${aws_vpc.yaum.id}"
}

resource "aws_subnet" "yaum" {
  availability_zone = "eu-west-1a"
  vpc_id                  = "${aws_vpc.yaum.id}"
  cidr_block              = "10.0.0.0/24"
  map_public_ip_on_launch = true

  depends_on = ["aws_internet_gateway.yaum"]
}


resource "aws_route_table_association" "a" {
  subnet_id      = "${aws_subnet.yaum.id}"
  route_table_id = "${aws_route_table.yaum.id}"
}

resource "aws_route_table" "yaum" {
  vpc_id = "${aws_vpc.yaum.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.yaum.id}"
  }
}

resource "aws_security_group" "yaum" {
  name        = "yaum"
  vpc_id = "${aws_vpc.yaum.id}"
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
        protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "yaum-01" {
  ami           = "ami-044b047d" // https://wiki.debian.org/Cloud/AmazonEC2Image/Stretch
  instance_type = "t2.micro"
  key_name = "${aws_key_pair.bjo.key_name}"
  tags {
    Name = "yaum-01"
  }
  subnet_id = "${aws_subnet.yaum.id}"
  vpc_security_group_ids = ["${aws_security_group.yaum.id}"]
}

resource "aws_eip" "yaum-01" {
  vpc = true
  instance                  = "${aws_instance.yaum-01.id}"
  depends_on                = ["aws_internet_gateway.yaum"]
}

resource "aws_eip_association" "yaum-01" {
  depends_on = ["aws_internet_gateway.yaum"]
  instance_id   = "${aws_instance.yaum-01.id}"
  allocation_id = "${aws_eip.yaum-01.id}"
}
