"""A Python Pulumi program"""

import pulumi
from pulumi_spacelift import Stack

stack = Stack("my-stack",
  administrative=False,
  autodeploy=False,
  branch="main",
  description="A simple stack",
  name="simple-stack-python-pulumi",
  repository="empty",
  project_root="",
  terraform_version="1.3.0"
)
