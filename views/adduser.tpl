{{template "layout.tpl"}}

<h1>Add New User</h1>

{{if .Success}}
<div class="alert alert-success">
    {{.Success}}
</div>
<script>
    setTimeout(function() {
        window.location.href = "/"; // 修改为你的首页URL
    }, 3000); // 3000毫秒后跳转，即3秒
</script>
{{else}}
<form action="/adduser" method="post" class="mt-4 mb-4">
    <div class="form-group">
        <label for="Username">Username:</label>
        <input type="text" class="form-control" name="Username" id="Username" required>
        {{if .Errors.Username}}
        <div class="alert alert-danger">{{.Errors.Username}}</div>
        {{end}}
    </div>
    <div class="form-group">
        <label for="Age">Age:</label>
        <input type="number" class="form-control" name="Age" id="Age" required>
        {{if .Errors.Age}}
        <div class="alert alert-danger">{{.Errors.Age}}</div>
        {{end}}
    </div>
    <div class="form-group">
        <label for="Email">Email:</label>
        <input type="email" class="form-control" name="Email" id="Email" required>
        {{if .Errors.Email}}
        <div class="alert alert-danger">{{.Errors.Email}}</div>
        {{end}}
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
{{end}}
