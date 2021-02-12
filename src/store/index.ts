import Vue from "vue";
import Vuex, { createLogger, ActionContext } from "vuex";
import { ActionTypes, InitializeAction, RefreshAction } from "./actions";
import { MutationTypes } from "./mutations";
import { query, StreamResponse } from "@/api";
import { Stream } from "./model";

export * from "./actions";
export * from "./mutations";
export * from "./model";

Vue.use(Vuex);

export class GlobalState {
    public streams: Stream[] = [];
    public version = 0;
}

type ActionParameters = ActionContext<GlobalState, GlobalState>;
/*
                if (state.version != reply.summary.version) {
                    commit(MutationTypes.VERSION, reply.summary.version);
                }
*/
export default new Vuex.Store<GlobalState>({
    plugins: [createLogger()],
    state: () => new GlobalState(),
    actions: {
        [ActionTypes.INITIALIZE]: async ({ dispatch, commit, state }: ActionParameters, payload: InitializeAction) => {
            await dispatch(new RefreshAction());
        },
        [ActionTypes.REFRESH]: async ({ dispatch, commit, state }: ActionParameters, payload: InitializeAction) => {
            await query<{ streams: StreamResponse[] }>("/v1/streams").then((reply) => {
                // console.log(`reply: ${JSON.stringify(reply)}`);
                commit(
                    MutationTypes.STREAMS,
                    reply.streams.map((sr) => new Stream(sr))
                );
            });
        },
    },
    mutations: {
        [MutationTypes.STREAMS]: (state: GlobalState, streams: Stream[]) => {
            Vue.set(state, "streams", streams);
        },
        [MutationTypes.VERSION]: (state: GlobalState, version: number) => {
            Vue.set(state, "version", version);
        },
    },
    modules: {},
});
