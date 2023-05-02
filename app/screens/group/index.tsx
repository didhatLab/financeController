import {GroupRouteProps, SpendGroup, GroupStack as Stack1, GroupStack} from "./types";
import {useState} from "react";
import {GroupContext} from "./context";
import {GroupsListScreen} from "./groups";
import {GroupSpendsScreen} from "./group-spends";
import {ViewGroupSpend} from "./edit-spending";
import {AddNewGroupSpendScreen} from "./add-group-spending";
import {MembersList} from "./members";
import {AdminGroup} from "./admin";
import {CreateNewGroup} from "./create-group";


export const GroupRoutes = (props: GroupRouteProps) => {
    const [groups, setGroups] = useState<SpendGroup[]>([])


    return (
        <GroupContext.Provider value={{groups: groups, setGroups: setGroups}}>
            <GroupStack.Navigator>
                <GroupStack.Screen name={'GroupList'} component={GroupsListScreen}/>
                <GroupStack.Screen name={'GroupScreen'} component={GroupSpendsScreen}/>
                <GroupStack.Screen name={'EditGroupSpend'} component={ViewGroupSpend}/>
                <GroupStack.Screen name={'AddNewGroupSpend'} component={AddNewGroupSpendScreen}/>
                <GroupStack.Screen name={'MembersList'} component={MembersList}/>
                <GroupStack.Screen name={'Admin'} component={AdminGroup}/>
                <GroupStack.Screen name={'CreateGroup'} component={CreateNewGroup}/>
            </GroupStack.Navigator>
        </GroupContext.Provider>
    )
}
