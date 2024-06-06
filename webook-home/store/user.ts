import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'
import $cookies from '~/utils/cookie'

export const state = () => ({
    tokenVal: '',
    visibleLoginDialog: false,
    info: {} as any,
})

export type UserModuleState = ReturnType<typeof state>

export const mutations: MutationTree<UserModuleState> = {
    UPDATE_TOKEN(state, payload: string) {
        state.tokenVal = payload
    },
    TOGGLE_LOGIN_DIALOG(state, payload: boolean) {
        state.visibleLoginDialog = payload
    },
    UPDATE_INFO(state, payload: any) {
        state.info = payload
    },
    UPDATE_AVATAR(state, payload: string) {
        state.info.avatar = payload
    },
    UPDATE_USERNAME(state, payload: string) {
        state.info.userName = payload
    },
}

export const getters: GetterTree<UserModuleState, RootState> = {}

export const actions: ActionTree<UserModuleState, RootState> = {
    async login({ commit, rootState }, token: string) {
        commit('UPDATE_TOKEN', token)
        $cookies.set(rootState.tokenKey, token, {
            maxAge: 60 * 60 * 24 * 7,
            path: '/',
        })
    },
    async logout({ commit, rootState }) {
        commit('UPDATE_TOKEN', '')
        commit('UPDATE_INFO', null)
        $cookies.remove(rootState.tokenKey)
    },
}
