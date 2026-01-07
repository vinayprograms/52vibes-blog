---
title: "Community"
description: "52vibes community projects. Add '52vibes' topic to your repository to join."
---

## Rules

1. **One week, one project** — Ship something real in 7 days
2. **Agent-first** — AI handles implementation; you steer
3. **No cherry-picking** — Document failures, not just wins

---

<div id="community-repos">
  <p class="loading">Loading repositories...</p>
</div>

<noscript>
  <p>Enable JavaScript to view community repositories, or visit the <a href="https://github.com/topics/52vibes">GitHub topic page</a>.</p>
</noscript>

<script>
(function(){
  var el=document.getElementById('community-repos'),
      cache=sessionStorage.getItem('52vibes-repos'),
      cacheTime=sessionStorage.getItem('52vibes-repos-time'),
      maxAge=60000;

  if(cache&&cacheTime&&Date.now()-parseInt(cacheTime)<maxAge){
    render(JSON.parse(cache));return;
  }

  fetch('https://api.github.com/search/repositories?q=topic:52vibes&sort=stars')
    .then(function(r){if(!r.ok)throw new Error('Rate limited');return r.json()})
    .then(function(d){
      sessionStorage.setItem('52vibes-repos',JSON.stringify(d.items||[]));
      sessionStorage.setItem('52vibes-repos-time',Date.now().toString());
      render(d.items||[]);
    })
    .catch(function(e){el.innerHTML='<p class="error">Could not load repositories. <a href="https://github.com/topics/52vibes">View on GitHub</a></p>'});

  function render(repos){
    if(!repos.length){el.innerHTML='<p class="empty">No community projects yet. Be the first!</p>';return;}
    var byOwner={};
    repos.forEach(function(r){
      var owner=r.owner.login;
      if(!byOwner[owner])byOwner[owner]=[];
      byOwner[owner].push(r);
    });
    var html='';
    Object.keys(byOwner).sort().forEach(function(owner){
      html+='<h3><a href="https://github.com/'+encodeURI(owner)+'">'+escapeHtml(owner)+'</a></h3>';
      html+='<ul class="repo-list">';
      byOwner[owner].forEach(function(r){
        var desc=r.description||'';
        if(desc.length>100)desc=desc.substring(0,97)+'...';
        html+='<li><a href="'+escapeHtml(r.html_url)+'">'+escapeHtml(r.name)+'</a> <span class="stars">★ '+String(r.stargazers_count|0)+'</span>';
        if(desc)html+='<p>'+escapeHtml(desc)+'</p>';
        html+='</li>';
      });
      html+='</ul>';
    });
    el.innerHTML=html;
  }

  function escapeHtml(s){
    var d=document.createElement('div');
    d.textContent=s;
    return d.innerHTML;
  }
})();
</script>
