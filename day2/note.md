## 组件化
1. 新建空文件夹
2. 创建并运行
    ```bash
    # 创建vite
    tyarn create vite
    # 对项目进行命名,默认名vite-project
    # 选择vue,然后选择vue-ts

    # 进入项目文件夹
    cd .\vite-project\
    # 初始化
    tyarn
    # 运行
    tyarn dev
    # 默认访问地址http://127.0.0.1:5173/
    ```
3. 对主要文件相互关系的理解
- index.html引入了main.ts文件
- main.ts引入了App.vue组件,并将其挂载到index.html
- 在网页审查元素时所看到的部分内容就来自于App.vue
4. 进一步理解App.vue文件
- style scoped标签
    - 表示该style只适用于当前文件,避免样式杂糅。
- template标签
    - template中的内容就是要挂载到index.html中的内容,能在网页审查中看见,template标签本身不可见
    - template标签虽然是空标签,但在组件化开发中用来规定哪些内容需要挂载到网页上,因此不能缺失
- HelloWorld标签
    - 这是一个自定义标签,配合js引入HelloWorld.vue组件
    - 该标签中的内容会传递给HelloWorld.vue组件,该标签中的参数名需和HelloWorld.vue组件中的参数名相对应
5. 开发练习
    1. 改写App.vue和HelloWorld.vue(两个组件的初始模板内容差不多)
    ```js
    <template>
      <div>这是一个组件模板</div>
    </template>
    
    <script lang="ts">
      // import引入
      import { defineComponent } from 'vue';
      import HelloWorld from './components/HelloWorld.vue';
      // export引出
      // 这里相当于const vm = new Vue({});
      export default defineComponent({
          components: {},
          data() {
          return {};
          },
          methods: {},
          computed: {},
      });
    </script>

    <style scoped></style>
    ```
    2. 修改HelloWorld.vue组件
    ```js
    <template>
      <div class="child">
       <h2>Hello World</h2>
       <div>{{ msg }}</div>
       <button @click="handleClick">Click me</button>
      </div>
    </template>

    <script lang="ts">
    import { defineComponent } from 'vue';
 
    export default defineComponent({
        components: {},
        data() {
          return {};
        },
        props: {
          // 接受来自App.vue的message
          msg: String,
        },
        methods: {
            handleClick() {
              // 不能直接在这里修改msg的值
              // 利用$emit能实现子组件对父组件的“逆向操作”
              // 可以通过$emit函数来修改父元素的值
              console.log(this);
              this.$emit("update:msg", "father");
              // 也可以通过$emit来绑定父元素的事件
              // 命名格式需要对应
              this.$emit("handleChildClick");
            },
        },
        computed: {},
    });
    </script>

    <style scoped>
    .child {
        margin-top: 150px;
        border: 1px solid black;
    }
    </style>
    ```
    3. 在App.vue中引入HelloWorld.vue组件
    ```js
    <template>
      <div>这是一个组件模板</div>
      <!-- 这里无法直接使用自定义标签,需要在components中声明 -->
      <input v-model="message"/>
      <!-- <HelloWorld :msg="message"/> -->
      <!-- 使用:msg可以实现动态文本,静态文本也可以直接用msg实现 -->
      <!-- 需要传递数字、布尔值、数组、对象时则需要借助v-bind -->
      <div>{{ message }}</div>

      <!-- 父组件(即App.vue)的值经过子组件的$emit修改后的值 -->
      <!-- <HelloWorld v-model:msg="message"/> -->

      <!-- 这里的@handle-child-click实际上是对应
      HelloWorld.vue中的$emit("handleChildClick")
      通过$emit子组件就能调用父组件的函数 -->
      <HelloWorld @handle-child-click="handleChildClick"/>
      <!-- 同理HelloWorld标签也可以写成hello-world -->
    </template>

    <script lang="ts">
    import { defineComponent } from 'vue';
    // 引入HelloWorld.vue组件
    import HelloWorld from './components/HelloWorld.vue';
  
    export default defineComponent({
        // 在components中声明组件
        components: { HelloWorld },
        data() {
          return {
            //传递message给HelloWorld.vue
            message: "hi",
          };
        },
        methods: {
            // 父组件点击事件
            handleChildClick() {
                console.log("child clicked");
            }
        },
        computed: {},
    });
    </script>

    <style scoped></style>
    ```
6. 补充说明
- App.vue默认为根组件,其他的组件都可以挂载到根组件上
- 导入数据时,一般选择根组件进行导入,这样就能直接对多个子组件分配数据;如果选择在子组件导入数据,则还需要通过$emit来将子组件的数据传递给父组件,再由父组件传递给其他子组件,增加了工作负担

