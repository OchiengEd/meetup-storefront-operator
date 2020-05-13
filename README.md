# Design of storefront
This is a simple operator created during the Kubernetes San Antonio Meetup based on operator-sdk.

## Type Store
Data we want to be able to define for the Store resource:
- Location: string
- Size: int 2

An example store for instance would appear as below:

```
apiVersion: app.storefront.com/v1alpha1
kind: Store
metadata:
  name: example-store
spec:
  location: "san antonio"
  size: 3
```

## GVK (Group Version Kind)
Storefront owns the domain `storefront.com`

Group: app.storefront.com  
Version: v1alpha1  
Kind: Store  
