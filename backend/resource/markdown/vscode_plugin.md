> VS Code 作为一个流行的编辑器，其轻便快速也使开发人员有很好的体验。我们在VS Code中安装的各式主题，Lint工具等，其实都是一个个的插件。VS Code内置了扩展能力，提供了一系列的插件API。基于此，我们可以很轻松地开发自己的插件，并应用到平时的开发中。



VS Code提供的API很多，本文通过实现一个简单的base64图片预览，来走通一个插件的开发发布使用流程
插件名称：byte-base64-viewer
插件使用方式：鼠标悬浮在base64之上，悬浮对应的图片，如下
![image.png](https://cdn.nlark.com/yuque/0/2021/png/716445/1626696199523-e4467fe0-2929-4e2e-93ee-1e57d20b4632.png#align=left&display=inline&height=95&margin=%5Bobject%20Object%5D&name=image.png&originHeight=189&originWidth=342&size=14337&status=done&style=none&width=171)

### 环境准备

1. Node 环境
1. VS Code编辑器



### 开发流程
#### 1. 安装插件项目生成器
```shell
npm install -g yo generator-code
```


#### 2. 初始化项目
```shell
yo code
```
随后跟进个人需求选择对应的选项，创建完成后进入用VS Code打开项目，如下
![image.png](https://cdn.nlark.com/yuque/0/2021/png/716445/1626695461321-5050b013-07f3-4481-aa93-68d02674e5c8.png#align=left&display=inline&height=267&margin=%5Bobject%20Object%5D&name=image.png&originHeight=534&originWidth=593&size=73319&status=done&style=none&width=296.5)
#### 3. 编码实现
初次开发插件只需注意两个文件，一个是 extension.ts，另一个是package.json
我们先看package.json
在VS Code中，插件都是懒加载的，所以我们需要设置激活的时机，package.json中的activationEvent字段即使激活的时机，其有多个激活时机可选择，具体可以查看官网。由于我们的插件是启动之后就要激活，将activationEvent改成下述方式:
```json
{
	...
	"activationEvents": ["*"],
  ...
 }
```
接下来看extension.ts
我们需要用到3个vscode提供的功能

1. vscode.languages.registerHoverPovider()：用来提供特定语言的鼠标Hover监听
1. vscode.Hover()：添加Hover事件
1. vscode.MarkdownString()：添加Markdown显示

代码如下
```typescript
export function activate(context: vscode.ExtensionContext) {
  const __hover = vscode.languages.registerHoverProvider(['javascript', 'typescript'], {
    provideHover: (document, position) => {
      const { _line } = position as any;
      const lineContent = document.lineAt(_line).text as any;
      const regexp = /('|")(data:image\/(jpeg|png|gif);base64,(.*))('|")/
			const res = lineContent.match(regexp);
      if (res) {
        const url = res[2];
        return new vscode.Hover(new vscode.MarkdownString(`![](${url})`));
      }
      return null;
    },
  });

  context.subscriptions.push(__hover);
}
```
activate钩子是VS Code提供的钩子，在插件激活时就会触发
至此，我们的编码就完成了。编码部分比较简单，如需更复杂的功能可查看官方文档的API


### 插件发布
#### 1. 注册publisher
首先去Microsoft的[Marketplace](https://marketplace.visualstudio.com/manage/publishers/louzhedong?auth_redirect=True)登录
注册一个publisher
![image.png](https://cdn.nlark.com/yuque/0/2021/png/716445/1626697405750-c45a88c1-53dc-4afa-b00c-f7c531ecafce.png#align=left&display=inline&height=164&margin=%5Bobject%20Object%5D&name=image.png&originHeight=654&originWidth=1438&size=71606&status=done&style=none&width=360)
#### 2. 打包插件
首先全局安装 vsce，一个用来打包发布的命令行工具
```shell
npm install -g vsce
```
随后执行打包命令
```shell
vsce package
```
打包完成后会生成一个vslx后缀的文件，将文件上传到Marketplace，几分钟之后，插件即发布完成，随后就可以在VS Code的插件市场搜索到
