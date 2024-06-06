declare module '@riophae/vue-treeselect' {
    import { Component } from 'vue'

    export const LOAD_ROOT_OPTIONS: string
    export const LOAD_CHILDREN_OPTIONS: string
    export const ASYNC_SEARCH: string
    export const VERSION: string

    const Treeselect: Component
    const treeselectMixin: Component
    export { Treeselect, treeselectMixin }
    export default Treeselect
}
