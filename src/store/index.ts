import Vue from "vue";
import Vuex, { createLogger, ActionContext } from "vuex";
import { ActionTypes, InitializeAction } from "./actions";
import { MutationTypes } from "./mutations";
import { query, StreamResponse } from "@/api";
import { Stream } from "./model";

Vue.use(Vuex);

export class GlobalState {
    public streams: Stream[] = [];
}

type ActionParameters = ActionContext<GlobalState, GlobalState>;

export default new Vuex.Store<GlobalState>({
    plugins: [createLogger()],
    state: () => new GlobalState(),
    actions: {
        [ActionTypes.INITIALIZE]: async ({ dispatch, commit, state }: ActionParameters, payload: InitializeAction) => {
            await query<{ streams: StreamResponse[] }>("/v1/streams").then((reply) => {
                console.log(`reply: ${JSON.stringify(reply)}`);
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
    },
    modules: {},
});
