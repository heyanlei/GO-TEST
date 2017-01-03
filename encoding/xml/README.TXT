我们看到Marshal函数接收的参数v是interface{}类型的，即它可以接受任意类型的参数，那么现在的问题是，在xml包，根据什么规则来生成相应的XML文件呢？

如果v是 array或者slice，那么输出每一个元素，类似value
如果v是指针，那么会Marshal指针指向的内容，如果指针为空，什么都不输出
如果v是interface，那么就处理interface所包含的数据
如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息
生成的XML文件中的element的名字又是根据什么决定的呢？元素名按照如下优先级从struct中获取：

如果v是struct，XMLName的tag中定义的名称
类型为xml.Name的名叫XMLName的字段的值
通过strcut中字段的tag来获取
通过strcut的字段名用来获取
marshall的类型名称
我们应如何设置struct 中字段的tag信息以控制最终xml文件的生成呢？
XMLName不会被输出
tag中含有"-"的字段不会输出
tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
tag中含有",chardata"，输出为xml的 character data而非element。
tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程
tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string
tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出

FirstName string   `xml:"name>first"`
LastName  string   `xml:"name>last"`
<name>
<first>Asta</first>
<last>Xie</last>
</name> 