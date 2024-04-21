<h1>Edit {{.Post.Title}}</h1>
<form action="" method="post">
    <div class="mb-3">
        <label for="title" class="form-label">标题:</label>
        <input type="text" class="form-control" id="title" name="title" value="{{.Post.Title}}">
    </div>
    <div class="mb-3">
        <label for="content" class="form-label">内容：</label>
        <textarea class="form-control" id="content" name="content" rows="10">{{.Post.Content}}</textarea>
    </div>
    <input type="hidden" name="id" value="{{.Post.Id}}">
    <button type="submit" class="btn btn-primary">Save Changes</button>
</form>
