<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>添加菜刀URL</title>
  <link rel="stylesheet" href="static/layui/css/layui.css">
</head>
<body>
<div style="width:700px;margin-top:20px;">
	<form class="layui-form" action="">
	  <div class="layui-form-item">
	    <label class="layui-form-label">URL</label>
	    <div class="layui-input-block">
      		<input type="text" name="url" required  lay-verify="required" placeholder="URL" autocomplete="off" class="layui-input">
	    </div>
	  </div>
	  <div class="layui-form-item">
	    <label class="layui-form-label">密码</label>
	    <div class="layui-input-block">
      		<input type="text" name="pwd" required  lay-verify="required" placeholder="密码" autocomplete="off" class="layui-input">
	    </div>
	  </div>
	  <div class="layui-form-item layui-form-text">
	    <label class="layui-form-label">文本域</label>
	    <div class="layui-input-block">
	      <textarea name="seting" placeholder="请输入内容" class="layui-textarea"></textarea>
	    </div>
	  </div>
	  <div class="layui-form-item">
	    <div class="layui-input-block">
	      <button class="layui-btn" lay-submit lay-filter="formDemo">立即提交</button>
	      <button type="reset" class="layui-btn layui-btn-primary">重置</button>
	    </div>
	  </div>
	</form>
</div>

<script type="text/javascript" src="static/js/jquery-3.0.0.min.js"></script>
<script type="text/javascript" src="static/layui/layui.js"></script>
<script>
layui.use('form', function(){
  var form = layui.form;
  //监听提交
  form.on('submit(formDemo)', function(data){
    $.ajax({
    	url :"/Add",
    	type:'POST',
			data:data.field,
    	success:function(msg){
				parent.reflash();
    	}
    })
    return false;
  });
});
</script>

</body>
</html>