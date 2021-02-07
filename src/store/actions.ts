export enum ActionTypes {
    INITIALIZE = "INITIALIZE",
    REFRESH = "REFRESH",
}

export class InitializeAction {
    type = ActionTypes.INITIALIZE;
}

export class RefreshAction {
    type = ActionTypes.REFRESH;
}
