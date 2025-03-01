package main

var projectWebTmpl = `
{{define "web_installation"}}
    <html>
    <head data-suburl="">
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <meta name="author" content="GoAdmin">
        <meta name="description"
              content="GoAdmin is a golang framework which helps gopher to build a data visualization and admin panel in ten minutes">
        <meta name="keywords" content="go, git, framework, admin, GoAdmin">

        <meta name="referrer" content="no-referrer">
        <meta name="_csrf" content="{{.CSRFToken}}">
        <meta name="_suburl" content="">

        <meta property="og:url" content="http://localhost:{{.Port}}/">
        <meta property="og:type" content="website">
        <meta property="og:title" content="GoAdmin">
        <meta property="og:description"
              content="GoAdmin is a golang framework which helps gopher to build a data visualization and admin panel in ten minutes.">
        <meta property="og:site_name" content="GoAdmin">

        <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
        <link rel="stylesheet" href="https://cdn.bootcdn.net/ajax/libs/font-awesome/4.6.3/css/font-awesome.min.css">
        <link rel="stylesheet" href="https://cdn.bootcdn.net/ajax/libs/octicons/8.5.0/build.css">
        <link rel="stylesheet" href="https://cdn.bootcdn.net/ajax/libs/semantic-ui/2.4.1/semantic.min.css">
        <noscript>
            <style>
                .dropdown:hover > .menu {
                    display: block;
                }

                .ui.secondary.menu .dropdown.item > .menu {
                    margin-top: 0;
                }
            </style>
        </noscript>

        <style>
            .install form .field {
                text-align: left;
            }

            h1, h2, h3, h4, h5, .ui.header, .ui.menu, .ui.input input, .ui.button:not(.label) {
                font-family: "PingFang SC", 'Hiragino Sans GB', "Helvetica Neue", "Microsoft YaHei", Arial, Helvetica, sans-serif !important;
            }

            .install form label {
                text-align: right;
                width: 320px !important;
            }

            .install form .field .help {
                margin-left: 335px !important;
            }

            .install form .field {
                text-align: left;
            }

            .form .help {
                color: #999999;
                padding-top: 0.6em;
                padding-bottom: 0.6em;
                display: inline-block;
                word-break: break-word;
            }

            .ui.attached.header {
                background: #f0f0f0;
            }

            .install {
                padding-top: 45px;
                padding-bottom: 80px;
            }

            body:not(.full-width) {
                font-family: "PingFang SC", "Helvetica Neue", "Microsoft YaHei", Arial, Helvetica, sans-serif !important;
                background-color: #fff;
                overflow-y: scroll;
                overflow-x: auto;
                min-width: 1020px;
            }

            .full.height {
                padding: 0;
                margin: 0 0 -80px 0;
                min-height: 100%;
            }

            .hide {
                display: none;
            }

            .install form input {
                width: 300px !important;
            }

            .install .ui.checkbox label {
                width: auto !important;
            }

            .install form .field.optional .title {
                margin-left: 320px !important;
            }

            .install .inline.checkbox {
                margin-top: -1em;
                margin-bottom: 2em;
            }

            .install .ui.checkbox {
                margin-left: 335px !important;
            }

            .ui.container:not(.fluid) {
                width: 980px !important;
            }

            .ui.left {
                float: left;
            }

            .ui.right {
                float: right;
            }

            footer .container .links > *:first-child {
                border-left: none;
            }

            footer .container .links > * {
                border-left: 1px solid #d6d6d6;
                padding-left: 8px;
                margin-left: 5px;
            }

            footer .container {
                padding-top: 10px;
            }

            footer {
                margin-top: 54px !important;
                height: 40px;
                background-color: white;
                border-top: 1px solid #d6d6d6;
                clear: both;
                width: 100%;
                color: #888888;
            }

            form .ui.dividing.header {
                background-color: #fbf8f8;
                color: #464545;
                padding-bottom: 6px;
                padding-top: 6px;
                margin-bottom: 25px;
            }
        </style>

        <script src="https://cdn.bootcdn.net/ajax/libs/semantic-ui/2.4.1/semantic.min.js"></script>

        <title>{{local "web.installation page"}} - GoAdmin</title>

        <meta name="theme-color" content="#65d4e8">
    </head>
    <body data-gr-c-s-loaded="true">
    <div class="full height">
        <div class="install">
            <div class="ui middle very relaxed page grid">
                <div class="sixteen wide center aligned centered column">
                    <h3 class="ui top attached header">
                        {{local "web.goadmin web installation program"}}
                    </h3>
                    <div class="ui attached segment install-form">
                        <form class="ui form" onsubmit="return false" action="##" method="post">

                            <h4 class="ui dividing header">{{local "web.database settings"}}</h4>
                            <div class="inline required field ">
                                <label>{{local "web.database type"}}</label>
                                <div class="ui selection database type dropdown" tabindex="0">
                                    <input type="hidden" id="db_type" name="db_type" value="MySQL">
                                    <div class="text">MySQL</div>
                                    <i class="dropdown icon"></i>
                                    <div class="menu" tabindex="-1">
                                        <div class="item active selected" data-value="mysql">MySQL</div>
                                        <div class="item" data-value="postgresql">PostgreSQL</div>
                                        <div class="item" data-value="mssql">MSSQL</div>
                                        <div class="item" data-value="sqlite">SQLite3</div>
                                        <div class="item" data-value="oceanbase">OceanBase</div>

                                    </div>
                                </div>
                            </div>

                            <div id="sql_settings" class="">
                                <div class="inline required field ">
                                    <label for="db_host">{{local "web.database host"}}</label>
                                    <input id="db_host" name="db_host" value="127.0.0.1">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_port">{{local "web.database port"}}</label>
                                    <input id="db_port" name="db_port" value="3306">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_user">{{local "web.database user"}}</label>
                                    <input id="db_user" name="db_user" value="root">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_passwd">{{local "web.database password"}}</label>
                                    <input id="db_passwd" name="db_passwd" type="password" value="">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_name">{{local "web.database name"}}</label>
                                    <input id="db_name" name="db_name" value="goadmin">
                                    <span class="help">{{local "web.where the framework sql data install to"}}</span>
                                </div>
                            </div>

                            <div id="pgsql_settings" class="hide">
                                <div class="inline required field ">
                                    <label for="db_schema">{{local "web.database schema"}}</label>
                                    <input id="db_schema" name="db_schema" value="public">
                                </div>
                            </div>

                            <div id="sqlite_settings" class="hide">
                                <div class="inline required field ">
                                    <label for="db_path">{{local "web.database file"}}</label>
                                    <input id="db_path" name="db_path" value="./admin.db">
                                    <span class="help">{{local "web.the file path of sqlite3 database"}}<br>{{local "web.please use absolute path when you start as service"}}</span>
                                </div>
                            </div>
                            <div id="oceanbase_settings" class="hide">
                                <div class="inline required field ">
                                    <label for="db_host">{{local "web.database host"}}</label>
                                    <input id="db_host" name="db_host" value="127.0.0.1">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_port">{{local "web.database port"}}</label>
                                    <input id="db_port" name="db_port" value="2883">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_user">{{local "web.database user"}}</label>
                                    <input id="db_user" name="db_user" value="root@sys">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_passwd">{{local "web.database password"}}</label>
                                    <input id="db_passwd" name="db_passwd" type="password" value="">
                                </div>
                                <div class="inline required field ">
                                    <label for="db_name">{{local "web.database name"}}</label>
                                    <input id="db_name" name="db_name" value="goadmin">
                                    <span class="help">{{local "web.where the framework sql data install to"}}</span>
                                </div>
                            </div>
                            <h4 class="ui dividing header">{{local "web.installation settings"}}</h4>

                            <div class="inline required field ">
                                <label>{{local "web.web framework"}}</label>
                                <div class="ui selection database type dropdown" tabindex="0">
                                    <input type="hidden" id="web_framework" name="web_framework" value="gin">
                                    <div class="text">Gin</div>
                                    <i class="dropdown icon"></i>
                                    <div class="menu" tabindex="-1">
                                        <div class="item active selected" data-value="gin">Gin</div>
                                        <div class="item" data-value="beego">Beego</div>
                                        <div class="item" data-value="buffalo">Buffalo</div>
                                        <div class="item" data-value="fasthttp">Fasthttp</div>
                                        <div class="item" data-value="echo">Echo</div>
                                        <div class="item" data-value="chi">Chi</div>
                                        <div class="item" data-value="gf">Gf</div>
                                        <div class="item" data-value="gorilla">Gorilla</div>
                                        <div class="item" data-value="iris">Iris</div>
                                    </div>
                                </div>
                            </div>
                            <div class="inline required field ">
                                <label for="module_name">{{local "web.module name"}}</label>
                                <input id="module_name" name="module_name" value="{{.DefModuleName}}" required="">
                                <span class="help">{{local "web.module name is the path of go module"}}</span>
                            </div>
                            <div class="inline required field ">
                                <label>{{local "web.theme"}}</label>
                                <div class="ui selection database type dropdown" tabindex="0">
                                    <input type="hidden" id="theme" name="theme" value="sword">
                                    <div class="text">Sword</div>
                                    <i class="dropdown icon"></i>
                                    <div class="menu" tabindex="-1">
                                        <div class="item active selected" data-value="sword">Sword</div>
                                        <div class="item" data-value="adminlte">Adminlte</div>
                                    </div>
                                </div>
                            </div>
                            <div class="inline required field ">
                                <label>{{local "web.language"}}</label>
                                <div class="ui selection database type dropdown" tabindex="0">
                                    <input type="hidden" id="language" name="language" value="cn">
                                    <div class="text">{{local "web.simplified chinese"}}</div>
                                    <i class="dropdown icon"></i>
                                    <div class="menu" tabindex="-1">
                                        <div class="item active selected"
                                             data-value="cn">{{local "web.simplified chinese"}}</div>
                                        <div class="item" data-value="en">{{local "web.english"}}</div>
                                        <div class="item" data-value="jp">{{local "web.japanese"}}</div>
                                        <div class="item" data-value="tc">{{local "web.traditional chinese"}}</div>
                                    </div>
                                </div>
                            </div>
                            <div class="inline required field ">
                                <label for="http_port">{{local "web.http port"}}</label>
                                <input id="http_port" name="http_port" value="80"
                                       required="">
                                <span class="help">{{local "web.port number which application will listen on"}}</span>
                            </div>
                            <div class="inline required field ">
                                <label for="prefix">{{local "web.url prefix"}}</label>
                                <input id="prefix" name="prefix" value="admin" required="">
                                <span class="help">{{local "web.url prefix of the running application"}}</span>
                            </div>
                            <div class="inline required field ">
                                <label>{{local "web.use orm"}}</label>
                                <div class="ui selection database type dropdown" tabindex="0">
                                    <input type="hidden" id="use_gorm" name="use_gorm" value="">
                                    <div class="text">{{local "web.no use"}}</div>
                                    <i class="dropdown icon"></i>
                                    <div class="menu" tabindex="-1">
                                        <div class="item active selected" data-value="">{{local "web.no use"}}</div>
                                        <div class="item" data-value="gorm">GORM</div>
                                    </div>
                                </div>
                            </div>

                            <h4 class="ui dividing header">{{local "web.application settings"}}</h4>

                            <div class="inline required field">
                                <label for="web_title">{{local "web.website title"}}</label>
                                <input id="web_title" name="web_title" value="GoAdmin"
                                       placeholder="{{local "input"}} {{local "web.website title"}}" required="">
                            </div>

                            <div class="inline required field">
                                <label for="login_page_logo">{{local "web.login page logo"}}</label>
                                <input id="login_page_logo" name="login_page_logo" value="GoAdmin"
                                       placeholder="{{local "input"}} {{local "web.login page logo"}}" required="">
                            </div>

                            <div class="inline required field">
                                <label for="sidebar_logo">{{local "web.sidebar logo"}}</label>
                                <input id="sidebar_logo" name="sidebar_logo" value="GoAdmin"
                                       placeholder="{{local "input"}} {{local "web.sidebar logo"}}" required="">
                            </div>

                            <div class="inline required field">
                                <label for="sidebar_min_logo">{{local "web.sidebar mini logo"}}</label>
                                <input id="sidebar_min_logo" name="sidebar_min_logo" value="GA"
                                       placeholder="{{local "input"}} {{local "web.sidebar mini logo"}}" required="">
                            </div>

                            <div class="ui divider"></div>
                            <div class="inline field">
                                <label></label>
                                <button class="ui primary button"
                                        style="background-color: #65d4e8;"
                                        onclick="install()">{{local "web.install now"}}</button>
                            </div>
                        </form>
                    </div>
                    <div class="ui attached segment install-success-readme-cn"
                         style="display:none;text-align: left;padding-left: 50px;padding-top: 30px;">
                        <h3>
                            一. 安装步骤
                        </h3>
                        <p>
                            1. 下载并导入对应的SQL文件：
                        </p>
                        <p>
                        <ul>
                            <li>
                                <a href="https://gitee.com/go-admin/go-admin/raw/master/data/admin.sql" target="_blank">Mysql</a>
                            </li>
                            <li>
                                <a href="https://gitee.com/go-admin/go-admin/raw/master/data/admin.mssql"
                                   target="_blank">Mssql</a>
                            </li>
                            <li>
                                <a href="https://gitee.com/go-admin/go-admin/raw/master/data/admin.pgsql"
                                   target="_blank">Postgresql</a>
                            </li>
                            <li>
                                <a href="https://gitee.com/go-admin/go-admin/raw/master/data/admin.db" target="_blank">SQLite3</a>
                            </li>
                            <li>
                                <a href="https://gitee.com/go-admin/go-admin/raw/master/data/admin.sql" target="_blank">OceanBase</a>
                            </li>
                        </ul>
                        <p>
                            2. 依次执行以下命令：
                        </p>
                        {{if eq .GOOS "windows"}}
                            <p>
                            <pre class="prettyprint lang-bsh">> GO111MODULE=on go mod init <span
                                        class="module_name_pre">xxx</span>
> GORPOXY=https://goproxy.io GO111MODULE=on go mod tidy
> GO111MODULE=on go run .</pre>
                            </p>
                        {{else}}
                            <p>
                            <pre class="prettyprint lang-bsh">> make init module=<span
                                        class="module_name_pre">xxx</span>
> GORPOXY=https://goproxy.io make install
> make serve</pre>
                            </p>
                        {{end}}
                        <p>
                            3. 登录访问
                        </p>
                        <p>
                            <a class="login_url_cn" href="http://127.0.0.1:9033/admin/login" target="_blank">http://127.0.0.1:9033/admin/login</a>
                        </p>
                        <p>
                            账号：admin&nbsp; 密码：admin
                        </p>
                        <p>
                            4. 查看文件中的readme了解更多
                        </p>
                        <h3>
                            二. 关于GoAdmin
                        </h3>
                        <h3>
                            <ul>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">Github：&nbsp;<a
                                                href="https://github.com/foundVanting/go-admin" target="_blank">https://github.com/foundVanting/go-admin</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">官网：<a
                                                href="https://www.go-admin.cn/"
                                                target="_blank">https://www.go-admin.cn/</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">论坛：<a
                                                href="http://discuss.go-admin.com/" target="_blank">http://discuss.go-admin.com/</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">文档：<a
                                                href="http://doc.go-admin.cn/zh/" target="_blank">http://doc.go-admin.cn/zh/</a></span>
                                </li>
                            </ul>
                        </h3>
                        <p>
                            <br/>
                        </p>
                    </div>
                    <div class="ui attached segment install-success-readme-en"
                         style="display:none;text-align: left;padding-left: 50px;padding-top: 30px;">
                        <h3>
                            一 Installation steps
                        </h3>
                        <p>
                            1. Download and import the corresponding SQL file:
                        </p>
                        <p>
                        <ul>
                            <li>
                                <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql"
                                   target="_blank">Mysql</a>
                            </li>
                            <li>
                                <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.mssql"
                                   target="_blank">Mssql</a>
                            </li>
                            <li>
                                <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.pgsql"
                                   target="_blank">Postgresql</a>
                            </li>
                            <li>
                                <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.db"
                                   target="_blank">SQLite3</a>
                            </li>
                            <li>
                                <a href="https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql"
                                   target="_blank">OceanBase</a>
                            </li>
                        </ul>
                        <p>
                            2. Execute the following command in turn
                        </p>
                        {{if eq .GOOS "windows"}}
                            <p>
                            <pre class="prettyprint lang-bsh">> GO111MODULE=on go mod init <span
                                        class="module_name_pre">xxx</span>
> GORPOXY=https://goproxy.io GO111MODULE=on go mod tidy
> GO111MODULE=on go run .</pre>
                            </p>
                        {{else}}
                            <p>
                            <pre class="prettyprint lang-bsh">> make init module=<span
                                        class="module_name_pre">xxx</span>
> GORPOXY=https://goproxy.io make install
> make serve</pre>
                            </p>
                        {{end}}
                        <p>
                            3. Login to visit
                        </p>
                        <p>
                            <a class="login_url_en" href="http://127.0.0.1:9033/admin/login" target="_blank">http://127.0.0.1:9033/admin/login</a>
                        </p>
                        <p>
                            Account: admin&nbsp; Password: admin
                        </p>
                        <p>
                            4. Learn more in the readme
                        </p>
                        <h3>
                            二. About GoAdmin
                        </h3>
                        <h3>
                            <ul>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">Github：&nbsp;<a
                                                href="https://github.com/foundVanting/go-admin" target="_blank">https://github.com/foundVanting/go-admin</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">Official Website: <a
                                                href="https://www.go-admin.com/"
                                                target="_blank">https://www.go-admin.com/</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">Forum: <a
                                                href="http://discuss.go-admin.com/" target="_blank">http://discuss.go-admin.com/</a></span>
                                </li>
                                <li>
                                    <span style="font-size:15px;font-weight:normal;">文档: <a
                                                href="http://doc.go-admin.cn/zh/" target="_blank">https://book.go-admin.cn/</a></span>
                                </li>
                            </ul>
                        </h3>
                        <p>
                            <br/>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="ui mini modal">
        <i class="close icon"></i>
        <div class="header">
            {{local "web.result"}}
        </div>
        <div class="content">

        </div>
        <div class="actions">
            <div class="ui approve button">{{local "web.ok"}}</div>
        </div>
    </div>
    <footer>
        <div class="ui container">
            <div class="ui left">
                © 2020 GoAdmin {{local "current version"}}: {{.Version}}
            </div>
            <div class="ui right links">
                <div class="ui language bottom floating slide up dropdown link item" tabindex="0">
                    <i class="world icon"></i>
                    <div class="text">{{local .CurrentLang}}</div>
                    <div class="menu" tabindex="-1">
                        <a class="item {{if eq .CurrentLang "web.simplified chinese"}}active selected{{end}}"
                           href="/?lang=cn">{{local "web.simplified chinese"}}</a>
                        <a class="item {{if eq .CurrentLang "web.english"}}active selected{{end}}"
                           href="/?lang=en">English</a>
                    </div>
                </div>
                <a href="http://localhost:5678/assets/librejs/librejs.html" style="display:none" data-jslicense="1">Javascript
                    Licenses</a>
                <a target="_blank" rel="noopener noreferrer"
                   href="https://www.go-admin.cn" style="color: #65d4e8;">{{local "web.official website"}}</a>
                <span class="version">{{.GoVer}}</span>
            </div>
        </div>
    </footer>

    <script>

        $("#db_type").change(function () {
            let sqliteDefault = './admin.db';

            let dbType = $(this).val();
            if (dbType === "sqlite") {
                $('#sql_settings').hide();
                $('#pgsql_settings').hide();
                $('#sqlite_settings').show();

                if (dbType === "sqlite") {
                    $('#db_path').val(sqliteDefault);
                }
                return;
            }

            let dbDefaults = {
                "mysql": {
                    "addr": "127.0.0.1",
                    "port": "3306",
                    "user": "root"
                },
                "postgresql": {
                    "addr": "127.0.0.1",
                    "port": "5432",
                    "user": "postgres"
                },
                "mssql": {
                    "addr": "127.0.0.1",
                    "port": "1433",
                    "user": "sa"
                },
                "oceanbase": {
                    "addr": "127.0.0.1",
                    "port": "2881",
                    "user": "root@sys"
                }
            };

            $('#sqlite_settings').hide();
            $('#sql_settings').show();
            $('#pgsql_settings').toggle(dbType === "postgresql");

            $.each(dbDefaults, function (key, value) {
                if ($('#db_type').val() === key) {
                    $('#db_host').val(value["addr"]);
                    $('#db_port').val(value["port"]);
                    $('#db_user').val(value["user"]);
                    return false;
                }
            });
        });


        $('#offline-mode input').change(function () {
            if ($(this).is(':checked')) {
                $('#disable-gravatar').checkbox('check');
                $('#federated-avatar-lookup').checkbox('uncheck');
            }
        });
        $('#disable-gravatar input').change(function () {
            if ($(this).is(':checked')) {
                $('#federated-avatar-lookup').checkbox('uncheck');
            } else {
                $('#offline-mode').checkbox('uncheck');
            }
        });
        $('#federated-avatar-lookup input').change(function () {
            if ($(this).is(':checked')) {
                $('#disable-gravatar').checkbox('uncheck');
                $('#offline-mode').checkbox('uncheck');
            }
        });
        $('#disable-registration input').change(function () {
            if ($(this).is(':checked')) {
                $('#enable-captcha').checkbox('uncheck');
            }
        });
        $('#enable-captcha input').change(function () {
            if ($(this).is(':checked')) {
                $('#disable-registration').checkbox('uncheck');
            }
        });

        function install() {
            let port = $('#http_port').val();
            let prefix = $('#prefix').val();
            let login_url = "http://127.0.0.1:" + port + "/" + prefix + "/login";

            $('.login_url_cn').attr("href", login_url);
            $('.login_url_cn').html(login_url);
            $('.login_url_en').attr("href", login_url);
            $('.login_url_en').html(login_url);

            let module_name = $('#module_name').val();
            $('.module_name_pre').each(function (index, ele) {
                $(ele).html(module_name);
            });

            $.ajax({
                type: "POST",
                dataType: "json",
                url: '/install?lang={{if eq .CurrentLang "web.simplified chinese"}}cn{{else}}en{{end}}',
                data: $('form.form').serialize(),
                success: function (data) {
                    console.log(data);
                    $('.ui.modal .content').html(data.msg);
                    $('.ui.modal').modal('show');
                    if (data.code === 0) {
                        $('.install-form').hide();
                        {{if eq .CurrentLang "web.simplified chinese"}}
                        $('.install-success-readme-cn').show();
                        {{else}}
                        $('.install-success-readme-en').show();
                        {{end}}
                    }
                },
                error: function () {
                    alert("error");
                }
            });
        }

        $('.ui.dropdown').dropdown({
            forceSelection: false
        });
        $('.ui.accordion').accordion();
        $('.ui.checkbox').checkbox();
        $('.poping.up').popup();

    </script>
    </body>
    </html>
{{end}}
`
