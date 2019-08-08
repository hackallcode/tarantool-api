s = box.schema.space.create('kv_storage')
s:format({{name = 'key', type = 'string'},{name = 'value'}})
s:create_index('primary', { type = 'hash', parts = {'key'} })