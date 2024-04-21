<h1>Blog posts</h1>

<div class="list-group">
{{range .blogs}}
    <a href="/view/{{.Id}}" class="list-group-item list-group-item-action flex-column align-items-start">
        <div class="d-flex w-100 justify-content-between">
            <h5 class="mb-1">{{.Title}}</h5>
            <small>{{formatTime .Created}}</small>
            
        </div>
        <p class="mb-1">Some summary text here...</p>
        <small>
            <a href="/edit/{{.Id}}" class="text-secondary">Edit</a> |
            <a href="/delete/{{.Id}}" class="text-danger">Delete</a>
        </small>
    </a>
{{end}}
</div>
