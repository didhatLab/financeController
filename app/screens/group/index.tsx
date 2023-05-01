import {GroupRouteProps, SpendGroup, GroupStack as Stack1, GroupStack} from "./types";
import {useState} from "react";
import {GroupContext} from "./context";
import {GroupsListScreen} from "./groups";
import {GroupSpendsScreen} from "./group-spends";
import {ViewGroupSpend} from "./edit-spending";


export const GroupRoutes = (props: GroupRouteProps) => {
    const [groups, setGroups] = useState<SpendGroup[]>([])


    return (
        <GroupContext.Provider value={{groups: groups, setGroups: setGroups}}>
            <GroupStack.Navigator>
                <GroupStack.Screen name={'GroupList'} component={GroupsListScreen}/>
                <GroupStack.Screen name={'GroupScreen'} component={GroupSpendsScreen}/>
                <GroupStack.Screen name={'EditGroupSpend'} component={ViewGroupSpend}/>
            </GroupStack.Navigator>
        </GroupContext.Provider>
    )
}
