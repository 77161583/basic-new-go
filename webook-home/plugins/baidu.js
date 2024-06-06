// var _hmt = _hmt || [];
// (function() {
// var hm = document.createElement("script");
// hm.src = "https://hm.baidu.com/hm.js?e0e427163a63bfb571e24c056aec361c";
// var s = document.getElementsByTagName("script")[0];
// s.parentNode.insertBefore(hm, s);
// })();

export default ({app: {router}, store}) => {
    /* 每次路由变更时进行pv统计 */
    router.afterEach((to, from) => {
      /* 告诉增加一个PV */
      try {
        window._hmt = window._hmt || []
        window._hmt.push(['_trackPageview', to.fullPath])
      } catch (e) {
      }
    })
}


