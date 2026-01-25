---
title: "Manual State Tracking Fixes Toc Highlighting"
weight: 23
agent: "Crush CLI Agent (model: claude-opus-4-5-20251101)"
duration: "2m 24s"
---

## TL;DR - Session insights

- IntersectionObserver's `isIntersecting` property only reflects instantaneous visibility - it doesn't maintain "current section" state, requiring manual tracking of which heading was last crossed
- Click-triggered scrolls don't reliably fire IntersectionObserver callbacks - explicit click handlers on TOC links must update highlighting state immediately
- The `rootMargin` configuration `-10% 0px -80% 0px` creates a narrow detection zone near the top of the viewport, ideal for determining which section the user is reading

---

## Session Log

* **USER**: Found a small issue with TOC entry highlighting script - The highlighting happens only when the section heading appears in the right area. If a scroll past a bit, even if the section heading is still visible (closer to the top than before), it stops highlighting it. Shouldn't the highlighting continue up till the next heading comes into that tracking area?

* **AI**: Looking at the TOC highlighting behavior.

* **TOOL-CALL**: grep `{"pattern":"IntersectionObserver","path":"./wk1-blog"}`

* **TOOL**:
  ```
  Found 13 matches
  ./wk1-blog/public/blog/week-01-blog-platform/index.html:
    Line 306, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/weeks/index.html:
    Line 140, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/blog/index.html:
    Line 118, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/about/index.html:
    Line 163, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/tags/hugo/index.html:
    Line 117, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/tags/week-1/index.html:
    Line 117, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/tags/infrastructure/index.html:
    Line 117, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/tags/experiment/index.html:
    Line 117, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/quarters/index.html:
    Line 108, Char 17:     var obs=new IntersectionObserver(function(entries){
  
  ./wk1-blog/public/tags/index.html:
    Line 132, Char 17:     var obs=new IntersectionObserver(function(entries){
  [...10 more lines...]
  ```

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/themes/52vibes/layouts/_default/baseof.html","offset":40,"limit":30}`

* **TOOL**:
  ```
  <file>
      41|    if(!toc)return;
      42|    var headings=document.querySelectorAll('.content h2[id], .content h3[id]');
      43|    if(!headings.length)return;
      44|    var links=toc.querySelectorAll('a[href^="#"]');
      45|    var linkMap={};
      46|    links.forEach(function(a){linkMap[a.getAttribute('href').slice(1)]=a});
      47|    var obs=new IntersectionObserver(function(entries){
      48|      entries.forEach(function(e){
      49|        var a=linkMap[e.target.id];
      50|        if(a)a.classList.toggle('active',e.isIntersecting);
      51|      });
      52|    },{root:document.querySelector('.content-area'),rootMargin:'-10% 0px -80% 0px',threshold:0});
      53|    headings.forEach(function(h){obs.observe(h)});
      54|  })();
      55|  </script>
      56|</body>
      57|</html>
  </file>
  
  ```

* **AI**: The issue is clear. The current logic only highlights when the heading is intersecting (visible in the detection zone). When you scroll past it, it stops highlighting. The fix is to track which section is "current" based on headings that have crossed the detection line.

* **TOOL-CALL**: edit `{"file_path":"./wk1-blog/themes/52vibes/layouts/_default/baseof.html","old_string":"    var obs=new IntersectionObserver(function(entries){\n      entries.forEach(function(e){\n        var a=linkMap[e.target.id];\n        if(a)a.classList.toggle('active',e.isIntersecting);\n      });\n    },{root:document.querySelector('.content-area'),rootMargin:'-10% 0px -80% 0px',threshold:0});\n    headings.forEach(function(h){obs.observe(h)});","new_string":"    var cur...`

* **TOOL**:
  ```
  <result>
  Content replaced in file: ./wk1-blog/themes/52vibes/layouts/_default/baseof.html
  </result>
  
  ```

* **TOOL-CALL**: bash `{"command":"cd ./wk1-blog && hugo --minify 2>&1 | head -20","description":"Rebuild Hugo site"}`

* **TOOL**:
  ```
  Start building sites … 
  hugo v0.154.2+extended+withdeploy darwin/arm64 BuildDate=2026-01-02T16:08:44Z VendorInfo=Homebrew
  
  
                    │ EN 
  ──────────────────┼────
   Pages            │ 25 
   Paginator pages  │  0 
   Non-page files   │  0 
   Static files     │  6 
   Processed images │  0 
   Aliases          │  0 
   Cleaned          │  0 
  
  Total in 117 ms
  
  
  <cwd>~/Documents/projects/52vibes</cwd>
  ```

* **AI**: Fixed. The TOC highlighting now tracks the "current" section and keeps it highlighted until the next heading enters the detection zone. When a heading scrolls out the top, it checks if another heading should take over as current.

* **USER**: Another bug... if I explicitly click on the section in TOC, the highlight doesn't change (even though the content scrolls correctly to put the heading on top of the page).

* **AI**: The issue is that when you click a TOC link, the page scrolls instantly but the IntersectionObserver doesn't fire reliably for click-triggered scrolls. Need to manually update the highlight on click.

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/themes/52vibes/layouts/_default/baseof.html","offset":40,"limit":20}`

* **TOOL**:
  ```
  <file>
      41|    if(!toc)return;
      42|    var headings=document.querySelectorAll('.content h2[id], .content h3[id]');
      43|    if(!headings.length)return;
      44|    var links=toc.querySelectorAll('a[href^="#"]');
      45|    var linkMap={};
      46|    links.forEach(function(a){linkMap[a.getAttribute('href').slice(1)]=a});
      47|    var current=null;
      48|    var obs=new IntersectionObserver(function(entries){
      49|      entries.forEach(function(e){
      50|        if(e.isIntersecting){current=e.target.id}
      51|        else if(e.boundingClientRect.top<0&&current===e.target.id){
      52|          var idx=[].indexOf.call(headings,e.target);
      53|          current=idx<headings.length-1?headings[idx+1].id:current;
      54|        }
      55|      });
      56|      links.forEach(function(a){a.classList.toggle('active',a.getAttribute('href')==='#'+current)});
      57|    },{root:document.querySelector('.content-area'),rootMargin:'-10% 0px -80% 0px',threshold:0});
      58|    headings.forEach(function(h){obs.observe(h)});
      59|  })();
      60|  </script>
  
  (File has more lines. Use 'offset' parameter to read beyond line 60)
  </file>
  
  ```

* **TOOL-CALL**: edit `{"file_path":"./wk1-blog/themes/52vibes/layouts/_default/baseof.html","old_string":"    var current=null;\n    var obs=new IntersectionObserver(function(entries){\n      entries.forEach(function(e){\n        if(e.isIntersecting){current=e.target.id}\n        else if(e.boundingClientRect.top<0&&current===e.target.id){\n          var idx=[].indexOf.call(headings,e.target);\n          current=idx<headings.length-1?headings[idx+1].id:current;\n        }\n      }...`

* **TOOL**:
  ```
  <result>
  Content replaced in file: ./wk1-blog/themes/52vibes/layouts/_default/baseof.html
  </result>
  
  ```

* **TOOL-CALL**: bash `{"command":"cd ./wk1-blog && hugo --minify 2>&1 | head -10","description":"Rebuild Hugo site"}`

* **TOOL**:
  ```
  Start building sites … 
  hugo v0.154.2+extended+withdeploy darwin/arm64 BuildDate=2026-01-02T16:08:44Z VendorInfo=Homebrew
  
  
                    │ EN 
  ──────────────────┼────
   Pages            │ 25 
   Paginator pages  │  0 
   Non-page files   │  0 
   Static files     │  6 
  
  
  <cwd>~/Documents/projects/52vibes</cwd>
  ```

* **AI**: Fixed. Added click handlers on TOC links that immediately set `current` to the clicked section and update the highlight.
