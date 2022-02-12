workspace {

    model {
        user = person "Personal User"
        softwareSystem = softwareSystem "Microservice System" "to handle user requests such as logins, signups, weekly reports, resources, etc." "Existing System" {
            postman = container "Postman" "Tools test API." "Postman" "Desktop App"
            backendProxy = container "Backend Proxy" "Backend proxy to split requests according to microservice" "Nginx" {
                user -> this "API calls"
                postman -> this "API calls"
            }
            msAuth = container "Microservice Auth" "Microservice Auth to manage login, signup" "Nodejs" {
                backendProxy -> this "Proxy API calls"
            }
            msFetch = container "Microservice Fetch" "Microservice Fetch to manage resource, resource USD, report weekly" "Golang" {
                backendProxy -> this "Proxy API calls"
            }
            database = container "Database File" {
                msAuth -> this "Reads from and writes to"
            }
        }
        
        eFisery = softwaresystem "Efisery" "The internal Efisery data resource." "Existing System"{
            msFetch -> this "Makes API calls and Reads Response" "JSON/HTTPS"
        }
        
        currConv = softwaresystem "currConv" "The internal currconv.com data Currency." "Existing System"{
            msFetch -> this "Makes API calls and Reads Response" "JSON/HTTPS"
        }
        
        
        deploymentEnvironment "Live" {
            deploymentNode "Postman" "" "Postman REST Tester" {
                postmanInstance = containerInstance postman
            }
            deploymentNode "Server Google" "" "Google Data Center" {
                deploymentNode "Docker Container" "" "Docker" {
                        proxyInstance = containerInstance backendProxy
                        authApiApplicationInstance = containerInstance msAuth
                        fetchApiApplicationInstance = containerInstance msFetch
                        jsonDatabaseInstance = containerInstance database
                }
            }
            deploymentNode "Server Efisery" "" "Efisery data center" "" {
                    softwareSystemInstance eFisery
            }
            
            deploymentNode "Server Currconv" "" "Currconv data center" "" {
                    softwareSystemInstance currConv
            }

        }
        
    }

    views {
        systemContext softwareSystem {
            include *
            autolayout lr
        }
        
        deployment softwareSystem "Live" "LiveDeployment" {
            include *
            autoLayout
        }

        styles {
            element "Existing System" {
                background #999999
                color #ffffff
            }
            element "Desktop App" {
                shape WebBrowser
            }
        }
        theme default
    }

}