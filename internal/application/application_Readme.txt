Application's Layer's Role:

- Defines primary ports (interfaces) like UserService
- It exposes methods that the driving adapter (e.g., an http controller) will call.

- Contains "use case logic" and the implementation of primary ports,
  encapsulate how the system behaves for specific actions (e.g. "create user").  