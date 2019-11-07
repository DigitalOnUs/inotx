### Experimental repo 
The idea is to import a json/hcl config with the mapping of load balancers per service
and translate it to a new picture with consul instead the LB. 

#### sample inputs 
-> samples/input.json

<div align="center">
#### sample usage
![sample](.img/sample.png)

input:
  ./inotx <initial_blue_print>.json|hcl
output:
  <initial_blue_print>consul.json|hcl

</div>

#### Pending 
Fix issues with the output json because hcl is not standard

WIP
