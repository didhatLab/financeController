import {createContext} from "react";
import {GroupContextType} from "./types";

export const defaultContext = {groups: [], setGroups: () => 0}
export const GroupContext = createContext<GroupContextType>(defaultContext)