output "out" {
  value = {
    lower : "${lower(var.in)}"
    upper : "${upper(var.in)}"
  }

  description = <<EOF
  An example output.
EOF
}
