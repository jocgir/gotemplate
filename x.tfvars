structure = {
  a = 1
  b = "hello"
  c = [1, 2, 3]

  d = {
    e = 5
  }

  f = {
    g = 10
  }
  h = upper("df")
  i = var.j
  k = var.l + 5 * 2
  s = "${var.l + 5 * 2}"
  s2 = "Hello ${var.l + 5 * 2}"
}

# This is a variable
variable "var1" {
  default     = 1
  description = "var1"
}

variable "var2" {
  default     = 2
  description = "var2"

  sub = {
    array = [1,
      2,
      "hello",
    ]
  }
}

resource "xxx" {}
