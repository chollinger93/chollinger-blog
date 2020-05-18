---
title: "Goodbye, Wordpress"
date: 2020-05-14T08:47:11+01:00
description: "Ditching Wordpress for a static site generator"
tags: ["linux", "web", "hugo", "html", "css", "go"]
draft: true
---

## Introduction 
For the past 4 or so years, I've been using Wordpress to push content to this blog. Maybe you've seen the sluggish mess that the blog is - and the fact that the server is hosted in Germany, whereas I am currently hosted in the United States, does not help with loading times.

<img>

## Workflow 
My **workflow**, however, was the real pain point (I accept judgment for this):

1. Write the content locally in Sublime or vim, ignoring formatting, so I don't have to deal with network issues and can work offline
2. Copy-paste the code to Wordpress, waiting for it to mess up
3. Manually add the code with an external SyntaxHighlighter plugin
4. Edit the generated HTML to fix any messed up code formatting
5. Cry and drink
6. Repeat this every 3 months

In other words, its worse than writing a Confluence page on dial up.

Accepting feedback and contributions is also difficult - while I have added feedback in the past, it requires me to do it. With a more standardized CI/CD process, it would be trivial to simply accept merge requests. 

But since the only person to blame here is me, I now finally managed to fix it.

# The Idea
> Using a static side generator is a really old idea and I don't claim any intellectual points for it. However, I am documenting this process for the sake of documenting it.

The idea here is to have a workflow that looks like this:
1. Write posts in markdown. Add code in pre-defined blocks/shortcodes.
2. Store images locally
3. Push to Git (GitHub, GitLab)
4. Deploy to the web server using a CI/CD flow

# Hugo
Hugo is a simple static-site generator written in `golang`. After thinking about `Jekyll` and the joys of Ruby, I've decided to go with the simpler alternative, as it only requires one binary, has a relatively simply syntax, a decent suite of tools and plugins, nice, open source / MIT licensend tempaltes and has been in my bookmarks since I saw it on HackerNews ages ago.

## Install Hugo
We just grab the latest binary for a 64-bit Linux from GitHub, unpack it, and store it somewhere on the `$PATH`.
{{< highlight zsh >}}
wget https://github.com/gohugoio/hugo/releases/download/v0.70.0/hugo_0.70.0_Linux-64bit.tar.gz
tar xvf hugo_0.70.0_Linux-64bit.tar.gz
{{< / highlight >}}

We can then set up a page:

{{< highlight zsh >}}
hugo new site chollinger-blog
{{< / highlight >}}


After some small edits to the `config.toml`, we should be good to go and only need content.

{{< highlight toml >}}
baseURL = "https://chollinger.com/blog"
languageCode = "en-us"
title = "Christian Hollinger"
theme = "ananke"
{{< / highlight >}}

## Exporting all old posts
I want to export all old blog posts to Markdown. There's a great little [tool](https://github.com/lonekorean/wordpress-export-to-markdown) by `lonekorean` on GitHub.

All we need is an export of Wordpress and run:

```
git clone https://github.com/lonekorean/wordpress-export-to-markdown
cd wordpress-export-to-markdown
npm install
node index.js
```

After that, we can copy the generated files to the posts directory.

## Testing

Start a dev server that automatically refreshes once we save a file:
{{< highlight zsh >}}
hugo server -D
{{< / highlight >}}

![First try](images/01-hugo.png)

Pretty? Well..


## Themes
We also need a real theme for the page.

My personal decision criteria for a theme were the following:
- No external dependencies to CDNs, googleapis, Google Analytics, trackers...
- Simplistic interface, without omitting information
- Fast load times on slow connections

I chose [hugo-ink](https://github.com/knadh/hugo-ink) by `knadh`, but customized it quite heavily to match some of my requirements.

{{< highlight zsh >}}
git clone https://github.com/knadh/hugo-ink.git themes/
git remote set-url origin https://github.com/otter-in-a-suit/hugo-ink.git
{{< / highlight >}}

Adjustments can be easily made using fairly standard HTML (w/ Hugos injections) and CSS, which I found easy to figure out, despite being anything but a Web Dev. :)

The **changes** made to that theme were the following:
- Removed all references to Google's font-CDN
- Removed Analytics code, even if it was controlled by a variable
- Modified the CSS to
  - Order all tags inline, as opposed to as a list
  - Change the background color for Syntax Highlighting, otherwise we're looking at grey code on a grey background
  - Added some classes for a Back button
- Added a Back button to all posts
- Added a TOC, controlled by a variable, to all posts
- Added a word count, tags, and an approximate read time to the overview
- Added very serious, random messages at the end of the posts

{{< highlight css >}}
.tag-li {
    display:inline !important;
}

.back {
    padding-top: 1em;
    font-size: 1em;
}
{{< / highlight >}}

The last part with the random messages was interesting, as `Hugo` / the underlying `go` Syntax allows you to pipe commands:
{{< highlight html "linenos=table" >}}
<div class="back">
	{{ if isset .Site.Params "footers" }}
		{{ if ne .Type "page" }}
			Next time, we'll talk about <i>"{{ range .Site.Params.footers | shuffle | first 1 }}{{ . }}"</i>{{ end }}
		{{ end }}
	{{ end }}
</div>
{{< / highlight >}}

And I am reasonably happy with the result:

![After adjustments](images/after.png)

## Adjusting exported blog posts
Due to my rather interesting `Wordpress` configuration, the exported posts from above need some help.

The issues I've found where the following:
- GitHub Gits are not rendered
- Internal Syntax Formatting caused everything to be escaped with `\`, breaking code
- Headlines are missing
- Tags are missing
- Descriptions are missing

### Gists
Gists get inserted as such:
![Export issues](images/issue3.png)

And look like that:
![Export issues](images/issue2.png)

Whereas we are expecting:
![Export issues](images/issue1.png)

We can fix that by replacing

{{< highlight md >}}
<script src="https://gist.github.com/otter-in-a-suit/8dcc79ab16daf5fcdda3df1a4ccc183a.js"></script>
{{< / highlight >}}

with

{{< highlight md >}}
{{</* gist otter-in-a-suit 8dcc79ab16daf5fcdda3df1a4ccc183a */>}}
{{< / highlight >}}

But, of course, doing that by hand would be tedious, so we can script that:

{{< highlight bash >}}
grep "gist.github.com" *.md | while read -r line ; do
    guser=$(echo "${line}" | cut -d'/' -f4)
    gist=$(echo "${line}" | cut -d'/' -f5 | sed -E 's/(\.js).+//g')
    echo "$gist"
    repl="{{/* gist "${guser}" "${gist}" */>}}"
    # Get line number
    ln=$(grep -n "$line" *.md | cut -d : -f 1)
    # Replace
    echo "Replace line $ln with ${repl}"
    sed -Ei "$ln s#.*#${repl}#g" *.md
done
{{< / highlight >}}

I've also added a `Visual Studio Code` shortcut in `keybindings.json` to insert Hugo's Syntax Highlighter blocks:

{{< highlight json "linenos=table" >}}

[
    {
        "key": "ctrl+1",
        "command": "editor.action.insertSnippet",
        "when": "editorTextFocus",
        "args": {
            "snippet": "{< highlight bash \"linenos=table\" >}"
        }
    },
    {
        "key": "ctrl+2",
        "command": "editor.action.insertSnippet",
        "when": "editorTextFocus",
        "args": {
            "snippet": "{< / highlight >}"
        }
    }
]{{< / highlight >}}

There were more adjustments that needed to be done - like re-adding videos - but for the most part, everything was working fine. I will spare you the tedious details.

## Deployment
For the deployment, the goal is to deploy the static HTML to the existing webserver over at [chollinger.com](https://chollinger.com), which runs `nginx`. Externally hosted sites, like `GitHub Pages`, are an option, but I would like to keep as much "in house" as possible.

We have 2 convinient (and free for Open Source) options for CI/build servers: `travis` and `GitHub Actions`. In this case, we'll be using `GitHub Actions`, as it avoids having yet another external dependency. 

Our flow will look like this:
1. Checkout the master branch
2. Update the theme via its submodule
3. Download hugo
4. Build the static HTML/CSS/JS
5. Deploy via SCP

Well create a `.github/workflows/workflow.yml` file as such:

{{< highlight yml "linenos=table" >}}
name: nginx
on: push
jobs:
  deploy:
    runs-on: ubuntu-18.04
    steps:
      - name: Git checkout
        uses: actions/checkout@v2
        
      - name: Update theme
        # Update themes
        run: git submodule update --init --recursive

      - name: Setup hugo
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: "0.70.0"

      - name: Build
        run: hugo --minify

      - name: Deploy
          uses: appleboy/scp-action@master
        with:
          host: ${{ secrets.HOST }}
          username: ${{ secrets.USERNAME }}
          passphrase: ${{ secrets.PASSWORD }}
          port: ${{ secrets.PORT }}
          key: ${{ secrets.KEY }}
          source: "public/*"
          target: "public"
{{< / highlight >}}

We'll deploy via SCP, but only with a specific user.

Add the user:
{{< highlight bash "linenos=table" >}}
useradd -m -d /home/github github
passwd github
su github
{{< / highlight >}}

Make the user's home owned by `root`:
{{< highlight bash "linenos=table" >}}
sudo chown -R root:github /home/github
sudo chown github:github /home/github/.ssh
# + Any folders the deployment needs
{{< / highlight >}}

Edit the SSH config:
{{< highlight bash "linenos=table" >}}
vim /etc/ssh/sshd_config
{{< / highlight >}}

Ensure this is all the user can do:
{{< highlight ini "linenos=table" >}}
Match User github
	ChrootDirectory %h
	ForceCommand /usr/lib/openssh/sftp-server -P read,remove
	AllowTcpForwarding no
{{< / highlight >}}

And restart the SSH daemon:
{{< highlight bash "linenos=table" >}}
service sshd restart
{{< / highlight >}}

Finally, change the user's shell to `/bin/true`:
{{< highlight bash "linenos=table" >}}
vim /etc/passwd
{{< / highlight >}}

## Conclusion
