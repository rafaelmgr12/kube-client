<h1 align="center">Kube Client</h1>
<p align = "center"> A Simple library to deal with API REST</p>

<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

The objective of this project is to develop a an HTTP client library for defined REST API. You can found this the problem definition in the following links
* [Challenge](https://app.devgym.com.br/challenges/9bcad7c4-a809-4ef5-929d-a000aede5b25)
* [API Documentaion](https://github.com/devgymbr/files/blob/main/kubeapi-docs.md)
* [How to run the API](https://github.com/devgymbr/files/blob/main/kubeapi-docs.md#como-rodar-a-api)

## 🚀 Technology

The library is built using 
* Go
* [UUID](https://github.com/google/uuid)

It provides a cohesive and straightforward interface to interact with a REST API, focusing on `Deployments` resource.

## 💻 Project

`kube-client` is an HTTP client library for a REST API that manages `Deployments` resources. It supports `POST`, `GET`, and `DELETE` operations and offers custom error handling for better user experience.

### Features:

- **Simple Client Initialization**: Easily create a new instance of the client with custom configurations.
- **CRUD Operations**: Perform `POST`, `GET`, and `DELETE` operations on `Deployments`.
- **Custom Error Handling**: Receive detailed error messages to understand and debug issues better.

## 📖 How to Run

### **Installation**:

```bash
   go get github.com/rafaelmgr12/kube-client
````

### Usage
* Initialize the Client:
```go
import "github.com/rafaelmgr12/kube-client/kube"

client := kube.NewClient(kube.WithBaseURL("http://api-url.com"), kube.WithTimeout(30))
```
* Create a Deployment
```go
deployment := &deployment.Deployment{
    Name:      "example-deployment",
    Namespace: "default",
    Replicas:  3,
}
createdDeployment, err := client.CreateDeployment(deployment)

```
* Retrieve a Deployment:
```go
deploymentID := "uuid-of-deployment"
retrievedDeployment, err := client.GetDeployment(deploymentID)


```
* Delete a Deployment:
```go
deploymentID := "uuid-of-deployment"
err := client.DeleteDeployment(deploymentID)


```

## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details
