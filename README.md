# go-rest-api
A RESTful API for a property management application.

Built with GoLang, mux, and MongoDB. 

## API 
__/property__
- `GET`: List all properties
- `POST`: Create a new property

__/property/{id}__
- `GET`: Get a property
- `PUT`: Update a property
- `DELETE`: Delete a property

__/tenant__
- `GET`: List all tenants
- `POST`: Create a new tenant

__/tenant/{id}__
- `GET`: Get a tenant
- `PUT`: Update a tenant
- `DELETE`: Delete a tenant

## TODO 
- [x] Define document models
- [x] Map out API Endpoints
- [x] Implement CRUD functionality
- [ ] Authenticate user HTTP requests 
- [ ] Write tests for APIS
- [ ] Write documentation
- [ ] Organize code in packages
- [ ] Build a deployment process