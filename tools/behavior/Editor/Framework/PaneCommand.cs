﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;

namespace Editor.Framework
{
    internal class PaneCommand : ICommand
    {
        #region fields

        readonly Action<object> _execute;
        readonly Predicate<object> _canExecute;

        #endregion fields

        #region Constructors

        public PaneCommand(Action<object> execute) : this(execute, null)
        {
        }

        public PaneCommand(Action<object> execute, Predicate<object> canExecute)
        {
            _execute = execute ?? throw new ArgumentNullException(nameof(execute));
            _canExecute = canExecute;
        }
        #endregion Constructors

        #region ICommand Members

        public bool CanExecute(object parameter)
        {
            return _canExecute?.Invoke(parameter) ?? true;
        }

        public event EventHandler? CanExecuteChanged
        {
            add { CommandManager.RequerySuggested += value; }
            remove { CommandManager.RequerySuggested -= value; }
        }

        public void Execute(object parameter)
        {
            _execute(parameter);
        }

        #endregion ICommand Members
    }
}
