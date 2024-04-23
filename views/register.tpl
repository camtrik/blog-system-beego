{{ template "layout.tpl" . }}

<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-6">
            {{ if .Success }}
                <div class="alert alert-success text-center" role="alert">
                    Registration successful! You will be redirected to login in <span id="countdown">5</span> seconds.
                </div>
                <script>
                    var seconds = 5;
                    function countdown() {
                        seconds -= 1;
                        document.getElementById('countdown').textContent = seconds;
                        if (seconds <= 0) {
                            window.location.href = "/login";
                        } else {
                            setTimeout(countdown, 1000);
                        }
                    }
                    countdown();
                </script>
            {{ else }}
                <h2>Register</h2>
                {{ if .GeneralError }}
                    <div class="alert alert-danger" role="alert">
                        {{.GeneralError}}
                    </div>
                {{ end }}
                <form action="/register" method="post">
                    <div class="mb-3">
                        <label for="username" class="form-label">Username</label>
                        <input type="text" class="form-control" id="username" name="username" value="{{.Username}}" required>
                        {{ if .UsernameErr }}
                            <div class="alert alert-danger" role="alert">{{.UsernameErr}}</div>
                        {{ end }}
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label">Password</label>
                        <input type="password" class="form-control" id="password" name="password" required>
                        {{ if .PasswordErr }}
                            <div class="alert alert-danger" role="alert">{{.PasswordErr}}</div>
                        {{ end }}
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label">Email</label>
                        <input type="email" class="form-control" id="email" name="email" value="{{.Email}}" required>
                        {{ if .EmailErr }}
                            <div class="alert alert-danger" role="alert">{{.EmailErr}}</div>
                        {{ end }}
                    </div>
                    <button type="submit" class="btn btn-primary">Register</button>
                </form>
            {{ end }}
        </div>
    </div>
</div>

{{ if .Success }}
<script>
    setTimeout(function() {
        window.location.href = "/login"; // Redirect to login page after 5 seconds
    }, 5000);
</script>
{{ end }}
