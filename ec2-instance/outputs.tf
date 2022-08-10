output "ubuntu_instance_public_ip" {
  description = "The IPv4 address of the Ubuntu instance. Enter this value into your RDP client when connecting to your instance."
  value       = "${aws_instance.ec2_instance.public_ip}"
}

output "instance_id" {
  value = "${aws_instance.ec2_instance.id}"
}