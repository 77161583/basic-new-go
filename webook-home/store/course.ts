import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { RootState } from '~/store'
import $cookies from '~/utils/cookie'

export const state = () => ({
    coursechosedata: {} as any
})

export type CourseModuleState = ReturnType<typeof state>
export const getters: GetterTree<CourseModuleState, RootState> = {
    coursechosedata: (state) => state.coursechosedata || {} as any,
}
export const actions: ActionTree<CourseModuleState, RootState> = {
    async chosedata({ commit, rootState }, coursechosedata: {}) {
        // console.log(coursechosedata,77777777)
        commit('CHOSEDATA', coursechosedata)
    }
}
export const mutations: MutationTree<CourseModuleState> = {
    CHOSEDATA(state, payload: string) {
        state.coursechosedata = payload
    },
}