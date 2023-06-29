
## Bitcask-go-tiny  
其项目主要是参考bitcask论文实现的K/V存储    


### 项目结构    
**索引**    
1.BTree构建索引Key和LogRecordPos （https://github.com/google/btree）   
2.自适应基数树 https://github.com/plar/go-adaptive-radix-tree 
3.B+树 https://github.com/etcd-io/bbolt  

**IO**  
1.flock（https://github.com/gofrs/flock） 


### 跑多个单元测试
例如，测试 fio/file_io_test.go中的所有方法 
```shell
go test -v ./
```


### refer
1.https://github.com/roseduan       
2.https://w02agegxg3.feishu.cn/docx/Ktp3dBGl9oHdbOxbjUWcGdSnn3g         
3.https://github.com/rfyiamcool     



