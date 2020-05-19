
To make this internet accessible on AWS, you may need an ALB per this doc:
https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html

Run the example commands listed in the above doc to:
1. Create an IAM OIDC provider
2. Create an IAM Policy
3. Apply an RBAC role via: 
'''
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/rbac-role.yaml
'''
4. Create an IAM Service Account
5. Apply an ALB Ingress Controller via: 
'''
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/alb-ingress-controller.yaml
'''
6. Edit the ALB ingress controller to add the cluster name
7. As long as the service type is set to "NodePort" for the other services,
   then deploy the ALB ingress yaml in this folder:
'''
kubectl apply -f buttercup_alb.yaml
'''


OLDER VERSIONS of AWS EKS may need the storage class setup per this doc:

https://docs.aws.amazon.com/eks/latest/userguide/storage-classes.html
