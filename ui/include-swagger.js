const fs = require('fs')
const path = require('path')
const yaml = require('js-yaml')
const swaggerYaml = fs.readFileSync(path.resolve(__dirname, '..', 'openapi.yml'), 'utf8')
const dfn = yaml.safeLoad(swaggerYaml)

const OUTPUT = path.join(__dirname, 'dist')


fs.writeFileSync(path.join(OUTPUT, 'openapi.json'), JSON.stringify(dfn, null, '\t'))
fs.writeFileSync(path.join(OUTPUT, 'openapi.yml'), swaggerYaml)
