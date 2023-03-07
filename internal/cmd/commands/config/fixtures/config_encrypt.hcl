# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

kms "aead" {
  purpose = "config"
  aead_type = "aes-gcm"
  key = "c964AJj8VW8w4hKz/Jd8MvuLt0kkcjVuFqMiMvTvvN8="
}

kms "aead" {
  purpose = "root"
  aead_type = "aes-gcm"
  key ="{{encrypt(eb78KqCwowELYnkOOko/XYz01q1ax3g76J1vCAvt5dQ=)}}"
}

kms "aead" {
  purpose = "worker"
  aead_type = "aes-gcm"
  key ="{{encrypt(aA1hxJo0JUAqcIATx/r0QTjAGD/btCPechEsukI2bt0=)}}"
}
