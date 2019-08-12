#!/usr/bin/env tarantool

box.cfg {
    listen = 3301
}
box.once("schema", function()
    s = box.schema.space.create('kv_storage')
    s:format({{name = 'key', type = 'string'},{name = 'value'}})
    s:create_index('primary', { type = 'hash', parts = {'key'} })
end)
