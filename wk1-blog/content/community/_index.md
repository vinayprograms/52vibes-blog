---
title: "Community"
description: "52vibes community projects and forks"
---

Projects and forks from the 52vibes community.

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
      maxAge=3600000;
  
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
    var html='<ul class="repo-list">';
    repos.forEach(function(r){
      var desc=r.description||'';
      if(desc.length>100)desc=desc.substring(0,97)+'...';
      html+='<li><a href="'+escapeHtml(r.html_url)+'">'+escapeHtml(r.full_name)+'</a>';
      html+='<span class="stars">★ '+r.stargazers_count+'</span>';
      if(desc)html+='<p>'+escapeHtml(desc)+'</p>';
      html+='</li>';
    });
    html+='</ul>';
    el.innerHTML=html;
  }
  
  function escapeHtml(s){
    var d=document.createElement('div');
    d.textContent=s;
    return d.innerHTML;
  }
})();
</script>
