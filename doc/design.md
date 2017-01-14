# Design

## Configuration

### Infrastructure
```
codename: r6
infra:
  - terra: git:infra_repo

provisioner:
  - name: chef
  - repo: ...
```


### Software deployment
```
deploy:
  - webfront
  - backend
```


## Tools

- Go-Fuzz for testing https://github.com/google/gofuzz