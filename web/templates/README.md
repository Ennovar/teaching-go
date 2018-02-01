# Templates

Templates in Go are supported by the package `html/template`. Templates are able to intake data and the data is automatically parsed to prevent code injection before being displayed to the page.

The upside to using templates as opposed to developing APIs is that it keeps your entire codebase very clean and prevents the need to write lots of JavaScript, not to say that you still might not need to write a little bit depending on the application/service.

The downside to using templates is that it makes development between frontend and backend teams very hard to accomplish. The reason that this is, is because the frontend code and backend code at this point is very much so intermingled. There is still a general separation, but not as much as there is when using APIs.