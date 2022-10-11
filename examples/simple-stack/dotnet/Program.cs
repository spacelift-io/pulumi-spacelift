using System.Collections.Generic;
using Pulumi;
using Pulumi.Spacelift;

using Stack = Pulumi.Spacelift.Stack;

return await Deployment.RunAsync(() =>
{
  // Add your resources here
  // e.g. var resource = new Resource("name", new ResourceArgs { });

  var myStack = new Stack("test", new StackArgs
  {
    Administrative = false,
    Autodeploy = false,
    Branch = "main",
    Description = "A simple stack",
    Name = "simple-stack-dotnet-pulumi",
    Repository = "empty",
    ProjectRoot = "",
    TerraformVersion = "1.3.0"
  });

  // Export outputs here
  return new Dictionary<string, object?>
  {
    ["stack_id"] = myStack.Id
  };
});
