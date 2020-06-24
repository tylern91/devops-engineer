# DevOps Engineer Assignment

## Tasks

1. Complete the Terraform stack to provide an EC2 instance in our AWS account
2. Write the main part for the role that would provision Doker to the instance
3. Write the main part for the role that would deploy the container as provided below
4. Solve the puzzle following the instructions inside the container.

## Requirements

### Terraform for EC2

The [Terraform](https://www.terraform.io) stack should be able to provide 1 EC2 instance with the following requirements:
* Add a SSH keypair to AWS from the terraform resources
* Use the latest **Ubuntu 20.04 Server** AMI, provided by Canonical
* Instance size: your choice (as long as it's able to handle the container processing)
* Attach the SSH keypair created above into the resource
* Please use latest Terraform version (prefer Terraform version >= 0.12)

**Notes**:
  * Provider and backend will be preconfigured (and will be overwritten when run). The code should resides in the file `terraform/main.tf`
  * We will not provide any AWS account/credential. Please create/use your own AWS account for the works.

### Docker provision

Write the [Ansible](https://docs.ansible.com/) role and its playbook which install the Docker components using the latest version available. This should reside in the file `ansible/roles/docker/tasks/main.yml`. Role prerequisite are on Ubuntu 20.04 server, as provisioned in Terraform.

Pull the latest tag of the Docker image and deploy to the server. Docker repository: `tylern91/devops-engineer`
Write the separate Ansible playbook or integrate the deployment step into the `ansible/appserver.yaml` - your call.

### Solve the puzzle

Once you finished the container deployment, making sure it's up and running. There is a simple puzzle hidden inside the container, try to solve it with your best effort to find out the easter egg. Write down your explanations on the way you solve it.

## Notes

All inventory are assumed and given at runtime - there is no need to manage it in the repo.
The playbook appserver.yml will be used to run the role and also the deployment part.
**You will be evaluate by following criterias:**
* Run-able Codes
* Results Verification (based on above requirements)
* Syntax issues
* Third-Party usage and referrences with explanations (reuse without ref to the sources in README is not acceptable and will be treated as fail)
* Code Quality (codes stucture; variables definitions; hardcoded ratio; best practice follow)
* Final answer of the puzzle mentioned above.

## Submit the task

Create your feature branch and work on this, then make the PR with the descriptions including all the configurations and explanations. Push to git.nfq.asia when done and notify career@nfq.asia as a final step.

For more information please contact quynhnhu.phan@nfq.asia
