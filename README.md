# Compiling wasm module
```shell
tinygo build -o process.wasm -target=wasi wasm.go   
```
# Environment Details
```
tinygo version 0.32.0 darwin/amd64 (using go version go1.22.4 and LLVM version 18.1.2)
go version go1.22.4 darwin/amd64
```
# Output
This code will run without any erros, but produces inconsistent output like:
```
Processed Value: fcacbfcf-0522-4a24-8689-a051fd088ae7 processed
Processed Value: 0ad0fb5c-d033-4c6b-a27a-87648664b29c processed
Processed Value: afc64584-a1b4-4b21-b259-260209d86c0c
Processed Value: 9bfc3f98-b91c-4365-b462-c24110f47ff2 processed
Processed Value: beffe719-5e5c-4f65-baff-99489dab3a01 processed
Processed Value: 0c6d2eea-3a81-48bd-96c9-cbfa06544367
Processed Value: 028da287-f6a4-49dc-b401-0d38e449882c processed
Processed Value: 7ba1e201-3fb0-48fe-b46b-3aba1357aee8 processed
Processed Value: 2abcd5d3-bd5e-443f-a436-c10f36b4b4df processed
Processed Value: b583873d-36a0-4ad4-b841-5cca72ddb0f5
Processed Value: 3186aca0-d726-456d-b002-eb855c862ba4 processed
Processed Value: 232eba55-7695-442a-bebb-8d4a2c877ac5 processed
Processed Value: ea4de08d-d754-4af3-8fe8-d226b0a4cc5e processed
Processed Value: c6d30f75-0bcc-4e4d-b95f-0f33a9a61a67 processed
Processed Value: a07abc59-ed0a-46bd-a752-d7956a991f20
Processed Value: e8ed76e1-0969-47f0-b2b5-7f084345cc80 processed
Processed Value: 235d47be-659f-490d-b1e0-f2b38bc05b4b processed
Processed Value: 7d5a14af-e680-4520-8dee-2939e70e4fdf processed
Processed Value: 66ef81aa-7544-4eeb-9ceb-981025ca8e1d processed
Processed Value: 2a5608a6-0d4d-4e5e-a921-7b3afddde8a8
Processed Value: 7dc8fe5e-800d-4bfd-860f-ea5f35980898 processed
Processed Value: 5a62b335-1fbe-451e-b238-1b2addabafdc processed
Processed Value: ab29c596-6697-42d7-bff7-693926a31dae processed
Processed Value: fbae4825-02d7-4f3d-b1ab-b31fa8f1c081 processed
Processed Value: 1736e140-97af-4618-b4db-c432c04b0b2f
Processed Value: e8561d84-b67a-4ab9-958b-d87de381dee5 processed
Processed Value: 0038efbb-4a48-4e18-aa55-e5cb279f2caf processed
Processed Value: d4d5c27c-ed64-4a49-886b-b5658ccd8e9b processed
Processed Value: 230a6968-acc7-4c0c-ac00-973cd868e74c processed
Processed Value: 16956c10-3d7b-4c32-b0a8-249d773bfe73 processed
Processed Value: 0700b7e5-9748-4185-a09f-bbe9abc070c6 processed
Processed Value: 8e32ea11-d6df-4a84-aa79-6714ab4b9833 processed
Processed Value: b514168b-76ad-410b-bcd1-a1dd5a6b2ebe
Processed Value: c51bba76-b2be-45a3-9e13-461ae3e201dd processed
Processed Value: 02f7a96e-1db6-4bae-a1e1-b0201e2077fb processed
Processed Value: cc5f1954-2a52-44a3-a649-7566cb3cfd50 processed
Processed Value: 6e0d61a5-e36d-42a9-8b1d-fd22a6532655 processed
Processed Value: 6eb54fdb-9c76-4989-b4a0-1e1c3306c6fc processed
Processed Value: 56f3a264-75cc-43f9-a5e9-8b7c21608ce5 processed
Processed Value: f39e67df-82a6-4042-8b43-449267a94d17 processed
Processed Value: 00fed8f2-73f7-42ff-8407-5b5f2d483c74 processed
Processed Value: ab5ab3f5-dd11-4991-ae33-9ebee676471a processed
Processed Value: be854539-b308-47c1-ae7b-028c3a105c96 processed
Processed Value: c956124b-3f51-41c9-8f80-e61f83c88c4b processed
Processed Value: 6378a668-6dd9-4466-8076-30440ccf25ef processed
Processed Value: aa55766d-c188-4dad-93e7-69d3825dacef processed
Processed Value: 2329ad3e-a228-48b3-8f66-a81f3ff55b14 processed
Processed Value: 9cc2b234-40a7-4d81-9717-ca90f881a3e8 processed
Processed Value: 16c27a1c-8e21-4afb-9222-6ef68e49b635 processed
Processed Value: 1f489dbe-5c49-4340-a141-a5c0ac976425 processed
```
