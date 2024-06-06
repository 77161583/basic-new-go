import Vue from 'vue'

const slider = Vue.directive('slider', {
    inserted(el, binding) {
        const defaultOptions = {
            mainCell: '.bd ul',
            autoPage: true,
            effect: 'leftLoop',
            autoPlay: false,
            scroll: 2,
            vis: 4
        }
        const customOptions: any = typeof binding.value === 'object' ? binding.value : null;
        (jQuery(el) as any).slide(customOptions || defaultOptions)
    }
})

export default slider
