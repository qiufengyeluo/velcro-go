﻿using Editor.Framework;
using Editor.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Editor.Commands
{
    class NewBehaviorTreeCommand : ViewModelCommand<EditorFrameViewModel>
    {
        public NewBehaviorTreeCommand(EditorFrameViewModel contextViewModel) : base(contextViewModel)
        {
        }

        public override void Execute(EditorFrameViewModel contextViewModel, object parameter)
        {
            if (contextViewModel.CurrWorkspace == null)
            {
                return;
            }

            string treeFileName = Utils.RandFileName.GetRandName(contextViewModel.CurrWorkspace.Dir,
                "NewBehaviorTree", ".json");
            if (string.IsNullOrEmpty(treeFileName) )
            {
                Dialogs.WhatDialog.ShowWhatMessage("错误", "创建行为树名失败");
                return;
            }

            var tr = new Datas.BehaviorTree() { FileName = treeFileName,
                Tree = new Datas.Files.Behavior3Tree() {
                    ID = Utils.ShortGuid.Next(),
                    Title= treeFileName.Replace(".json", ""),
                    Description = ""
                }
            };
            contextViewModel.CurrWorkspace.Trees.Add(tr);
            if (!contextViewModel.IsWorkspaceExpanded)
            {
                contextViewModel.IsWorkspaceExpanded = true;
            }
        }

        public override bool CanExecute(EditorFrameViewModel contextViewModel, object parameter)
        {
            if (contextViewModel.IsReadOnly)
            {
                return false;
            }
            return true;
        }
    }
}