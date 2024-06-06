let text = '';
(function(a, h, c, b, f, g) { a.UdeskApiObject = f; a[f] = a[f] || function() { (a[f].d = a[f].d || []).push(arguments) }; g = h.createElement(c); g.async = 1; g.charset = 'utf-8'; g.src = b; c = h.getElementsByTagName(c)[0]; c.parentNode.insertBefore(g, c) })(window, document, 'script', 'https://assets-cli.s4.udesk.cn/im_client/js/udeskApi.js', 'ud')
ud({
  code: '2jhbcbc',
  link: 'https://clkefu.s4.udesk.cn/im_client/?web_plugin_id=4296'
})

function showKefu() {
  $('#udesk_btn_text').trigger('click')
}
text += '<a class="right-online" href="javascript:void(0);" onclick="showKefu()">'
text += '<img src="../images/sczj/icon-kf.png"'
text += '<h6>在线咨询</h6>'
text += '</a>'
document.write(text)
document.write('<style>#udesk_btn{display: none;}</style>')

function appendText() {
  const txt1 = ''
  const txt2 = ''
  const txtTj = ''
  const txt3 = document.createElement('p')
  txt3.innerHTML = 'Text.' // 閫氳繃 DOM 鏉ュ垱寤烘枃??
  // 杩藉姞鏂板厓??
  const locations = window.location.host
  const arr = ['xamzj.chinahrt.com', 'hnzj.chinahrt.com', 'hangyeke.chinahrt.com', 'xcwjw.chinahrt.com', 'nmghjgc.chinahrt.com', 'gszj.chinahrt.com', 'qingshuihe.chinahrt.com',
    'nmg.chinahrt.com', 'helin.chinahrt.com', 'tlzj.chinahrt.com', 'nmgjztp.chinahrt.com', 'baotouzj.chinahrt.com', 'sanya.chinahrt.com', 'saihan.chinahrt.com',
    'dljxjy.ncepu.edu.cn', 'nmggwy.chinahrt.com', 'hkzj.chinahrt.com', 'nmgcf.chinahrt.com', 'ts.chinahrt.com', 'gqb.chinahrt.com']
  let status = 0
  if (locations === 'tjjxjy.chinahrt.com') {
    status = 3
  } else {
    for (let i = 0; i < arr.length; i++) {
      // console.log($("iframe").contents().find("#platformId").val(),'ppppp');
      if (arr[i] === locations) {
        status = 1
      }
    }
  }
  if (status === 1) {
    $('body').append(txt2)
  } else if (status === 3) {
    $('body').append(txtTj)
  } else {
    $('body').append(txt1)
  }
}
function closeDiv() {
  document.getElementById('adfloat_fy').style.display = 'none'
}

$(function() {
  appendText()
})
