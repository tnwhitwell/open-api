const SwaggerUI = require('swagger-ui')

SwaggerUI({
  dom_id: '#container',
  url: '/openapi.json',
  displayOperationId: true,
  deepLinking: true
})
