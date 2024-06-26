# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

function has_authorized_action() {
  # accepts the output of a read on a arbitrary resource that has authorized_actions in its
  # output and the action to expect in the list as the second argument:
  #    has_authorized_action $out authorize-session
  local out=$1
  local action=$2
  echo $out | jq -c ".item.authorized_actions | contains([\"$action\"])"
}
