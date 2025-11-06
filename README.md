# Terraform Provider ConfDB

***work in progress***

This is a custom Terraform Provider based on HashiCorp's [Implement a provider with the Terraform Plugin Framework](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider) guide.

Basic premise of this project is that, instead of using libraries and configuration facilities built into CI/CDs (For example Azure Devops ["variable groups"](https://learn.microsoft.com/en-us/azure/devops/pipelines/library/)), you can use Terraform-native way.

Some of the benefits of this approach:
- Portability (you can use this provider from your local machine, other CI/CD etc. - from where you are using Terraform).
- Version control (Frequently, such inputs are not covered by same guarantees, policies as your code; it's often more difficult to structure it and notice rot or changes).
- Ease of use (it's potentially easier to work with Terraform provider than to wrestle with CI/CD-related tools, DSLs, UIs etc.)
