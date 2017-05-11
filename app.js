"use strict";

var env = process.env.NODE_ENV = process.env.NODE_ENV || 'development';
var port = process.env.PORT || 8099;

var argv = require('optimist')
	.usage('launch the risp-search web service.\nUsage: $0 [-pmr] [--init]')
	.default('p', port)
	.alias('p', 'port')
	.describe('p', 'port to host search features on')
	.default('m', './config/RISP_SearchLocalMongos.json')
	.alias('m', 'mongo')
	.describe('m', 'options file defining mongo connection. defaults to ./config/RISP_SearchLocalMongos.json')
	.default('r', './config/RISP_SearchLocalRedis.json')
	.alias('r', 'redis')
	.describe('r', 'options file defining redis connection. defaults to ./config/RISP_SearchLocalRedis.json')
	.boolean('i')
	.alias('i', 'init')
	.describe('i', 'initialize redis with a fresh copy of LocationData before starting')
	.boolean('d')
	.alias('d', 'drop')
	.describe('d', 'drop the existing data-set prior to loading')
	.default('s', 'IMLS')
	.alias('s', 'source')
	.describe('s', 'MLS resource used for prefixing location data')
	.alias('u', 'Users')
	.describe('u', 'Sync User collection')
	.default('j', false)
	.alias('j', 'json')
	.describe('j', 'import locations from a JSON file')
	.argv;