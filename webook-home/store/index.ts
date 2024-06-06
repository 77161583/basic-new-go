import { GetterTree, ActionTree, MutationTree } from 'vuex'

export const state = () => ({
    tokenKey: 'clzjToken',
})

export type RootState = ReturnType<typeof state>

export const getters: GetterTree<RootState, RootState> = {
    isLogin: (state) => !!state['user'].tokenVal,
    userInfo: (state) => state['user'].info || {},
}

export const mutations: MutationTree<RootState> = {}

export const actions: ActionTree<RootState, RootState> = {
    nuxtServerInit({ commit, state }, { app }) {
        const token = app.$cookies.get(state.tokenKey) || ''
        commit('user/UPDATE_TOKEN', token)
    },
}
